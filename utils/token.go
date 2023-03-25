package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	profileServices "event-backend/service/profileServices"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// TokenBlacklist is a map of blacklisted tokens, to keep track of invalid tokens ---------------
var TokenBlacklist = make(map[string]bool)

func AddToBlacklist(tokenString string) {
	TokenBlacklist[tokenString] = true
}

func RemoveFromBlacklist(tokenString string) {
	delete(TokenBlacklist, tokenString)
}

func ClearTokenBlacklist() {
	TokenBlacklist = make(map[string]bool)
}
// TokenWhitelist is a map of valid tokens to keep track of which tokens are have access rights. 
// In case any user should be logged out, token in this list can be taken into blacklist -------
var TokenWhitelist = make(map[string]bool)

func AddToWhitelist(tokenString string) {
	TokenWhitelist[tokenString] = true
}

func RemoveFromWhitelist(tokenString string) {
	delete(TokenWhitelist, tokenString)
}

func ClearTokenWhitelist() {
	TokenWhitelist = make(map[string]bool)
}
//---------------------------------------------------------------------------------------------

// function to check whether the token is valid or not
func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

// function to generate token for the user with the given credentials
func GenerateToken(userId uint, role string) (string, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	profile, err := profileServices.GetByUserId(userId)
	if err != nil {
		return "", err
	}
	claims["profileId"] = profile.ID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	AddToWhitelist(tokenString)

	return tokenString, err
}

// function to extract token from the request header
func ExtractToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}

	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// function to extract user id from the token
func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["profileId"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}
