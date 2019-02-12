package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var projectPath string

func init() {
	// init config
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	p := viper.GetStringMapString("project")
	projectPath = p["name"]
	initDatabase()
	initServer()
}

func main() {
}
