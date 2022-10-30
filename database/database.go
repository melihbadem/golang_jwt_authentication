package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jwt-authentication/models"
	"os"
)

var DBConn *gorm.DB

func Main() {
	var err error
	DBConn, err = gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DBConn.AutoMigrate(models.User{})
	fmt.Print("database connection made successfully")
}
