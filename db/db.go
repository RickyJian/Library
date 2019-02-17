package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"library/config"
	"time"
)

var connGlob *gorm.DB
var IsMigrate = false

func init() {
	db := config.GetDB()
	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", db.Host, db.Port, db.User, db.DBName, db.Password, db.SSLMode)
	database, err := gorm.Open("postgres", dbConn)
	database.DB().SetMaxOpenConns(10)
	database.DB().SetMaxIdleConns(5)
	database.DB().SetConnMaxLifetime(time.Second * 300)
	connGlob = database
	IsMigrate = db.IsMigrate
	if err != nil {
		fmt.Println(err)
	}
	database.SingularTable(true)
}

func GetConn() *gorm.DB {
	return connGlob
}
