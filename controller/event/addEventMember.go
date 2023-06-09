package event

import (
	"event-backend/model"
	groupServices "event-backend/service/group"
	groupMemberServices "event-backend/service/groupMember"
	profileServices "event-backend/service/profile"
	"event-backend/util"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	lockMap      = make(map[string]*sync.Mutex)
	lockMapMutex = sync.Mutex{}
)
// also some kind of message queue can be used for concurrency instead of mutex locks
func getLock(eventId string) *sync.Mutex {
	// Use a map to store mutex locks per event ID
	// This ensures that each event has its own mutex lock
	lockMapMutex.Lock()
	defer lockMapMutex.Unlock()

	lock, ok := lockMap[eventId]
	if !ok {
		lock = &sync.Mutex{}
		lockMap[eventId] = lock
	} else {
		if lock == nil {
			lock = &sync.Mutex{}
			lockMap[eventId] = lock
		}
	}

	return lock
}

// in addition to its concurrent structure, it also can work more efficiently by using goroutines and gochannels as asycnchronous tasks
func AddEventMemberController(c *gin.Context) {
	var group *model.Group
	var groupMember model.GroupMember

	// Extract player ID from JWT authorization token
	playerId, err := util.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error!", "data": err.Error()})
		return
	}

	// Get event ID from request parameters
	eventID_str := c.Query("eventId")

	// Obtain the mutex lock
	lock := getLock(eventID_str)

	// Lock the mutex to ensure exclusive access to the code below
	lock.Lock()
	defer lock.Unlock()

	// convert string eventId to uint
	eventId, err := strconv.ParseUint(eventID_str, 10, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error!", "data": err.Error()})
		return
	}

	// Get player profile by ID we extracted from the authorization token
	playerProfile, err := profileServices.GetById(playerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error on get by id!", "data": err.Error()})
		return
	}

	// Check if player already joined the event, this method uses transaction to protect the data integrity
	if groupMemberServices.HasJoinedEvent(playerId, eventId) {
		// If the user has already joined an event, return the group ID and the list of player IDs in the group as the registration is successful
		group, err = groupServices.GetByPlayerId(playerId, uint(eventId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error on get player by id!", "data": err.Error()})
			return
		}
		players, err := groupServices.GetGroupMembers(group.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error on get group members!", "data": err.Error()})
			return
		}
		// get list of ids of the players in the group
		playerIds := make([]uint, len(players))
		for i, player := range players {
			playerIds[i] = player.ID
		}
		c.JSON(http.StatusOK, gin.H{"message": "player joined event successfully", "groupID": group.ID, "playerIds": playerIds})
		return
	}

	// Get group by level, this method uses transaction to protect the data integrity
	group, err = groupServices.GetByLevel(playerProfile.Level, eventId)
	if err != nil {
		// if there is no such group for the player's level, create a INITIAL group
		if err.Error() == "record not found" {
			// Create new group
			newGroup := model.Group{
				EventID: uint(eventId),
				// GroupCategoryID is hardcoded since the document says so but in the future it can be dynamic since the model and database already support it
				GroupCategoryID: func(playerProfile *model.Profile) uint {
					if playerProfile.Level >= 0 && playerProfile.Level < 20 {
						return 1
					} else if playerProfile.Level >= 20 && playerProfile.Level < 50 {
						return 2
					} else {
						return 3
					}
				}(playerProfile),
				MaxMember: 20,
			}
			// Create new group
			group, err = groupServices.Create(&newGroup)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error on create!", "data": err.Error()})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error on get by level!", "data": err.Error()})
			return
		}
	}

	// Check if group is full, this method uses transaction to protect the data integrity
	if groupServices.IsGroupFull(group.ID) {
		// Create new group
		newGroup := model.Group{
			EventID: uint(eventId),
			// GroupCategoryID is hardcoded since the document says so but in the future it can be dynamic since the model and database already support it
			GroupCategoryID: func(playerProfile *model.Profile) uint {
				if playerProfile.Level >= 0 && playerProfile.Level < 20 {
					return 1
				} else if playerProfile.Level >= 20 && playerProfile.Level < 50 {
					return 2
				} else {
					return 3
				}
			}(playerProfile),
			MaxMember: 20,
		}
		// Create new group
		group, err = groupServices.Create(&newGroup)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error!", "data": err.Error()})
			return
		}
	}

	// Add player to group
	groupMember = model.GroupMember{
		GroupID:   group.ID,
		ProfileID: playerProfile.ID,
	}
	_, err = groupMemberServices.Create(&groupMember)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error!", "data": err.Error()})
		return
	}

	players, err := groupServices.GetGroupMembers(group.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error!", "data": err.Error()})
		return
	}
	// get list of ids of the players in the group
	playerIds := make([]uint, len(players))
	for i, player := range players {
		playerIds[i] = player.ID
	}
	c.JSON(http.StatusOK, gin.H{"message": "player joined event successfully", "groupID": group.ID, "playerIDs": playerIds})
}
