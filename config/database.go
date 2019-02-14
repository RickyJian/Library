package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"library/db"
)

func InitDatabase(){
	dbConfig := viper.GetStringMapString("database")
	host := dbConfig["host"]
	port := dbConfig["port"]
	user := dbConfig["user"]
	dbName := dbConfig["name"]
	password := dbConfig["password"]
	sslmode := dbConfig["sslmode"]
	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",host,port,user,dbName,password,sslmode)
	database, err := gorm.Open("postgres", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	database.SingularTable(true)
	database.AutoMigrate(&db.Book{})
	defer database.Close()
}