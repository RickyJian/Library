package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Server struct {
	Root string
}

type Database struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func GetProject() (s Server) {
	p := viper.GetStringMapString("project")
	s = Server{Root: p["name"]}
	return
}

func GetDB() (db Database) {
	dbConfig := viper.GetStringMapString("database")
	host := dbConfig["host"]
	port := dbConfig["port"]
	user := dbConfig["user"]
	dbName := dbConfig["name"]
	password := dbConfig["password"]
	sslmode := dbConfig["sslmode"]
	db = Database{host, port, user, dbName, password, sslmode}
	return
}
