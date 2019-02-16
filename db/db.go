package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"library/config"
)

var connGlob *gorm.DB

func init() {
	db := config.GetDB()
	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", db.Host, db.Port, db.User, db.DBName, db.Password, db.SSLMode)
	database, err := gorm.Open("postgres", dbConn)
	connGlob = database
	if err != nil {
		fmt.Println(err)
	}
	database.SingularTable(true)
	//database.AutoMigrate(&repository.Book{})

}

func GetConn() *gorm.DB {
	return connGlob
}
