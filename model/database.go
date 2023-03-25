package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDataBase() {
	Dbdriver := os.Getenv("DIALECT")
	DbHost := os.Getenv("HOST")
	DbUser := os.Getenv("USER")
	DbPassword := os.Getenv("PASSWORD")
	DbName := os.Getenv("DBNAME")
	DbPort := os.Getenv("PORT")

	DBURL := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	var err error
	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("connected to database ", DbName)
	}

	// DB.AutoMigrate(&User{})
	
	// DB.AutoMigrate(&Profile{})
	// DB.AutoMigrate(&Profile{}, &User{})
	// DB.Model(&Profile{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

}

func DisconnectDataBase() {
	DB.Close()
}
