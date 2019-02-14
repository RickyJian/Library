package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"library/config"
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
	config.InitDatabase()
	s := config.Server{Root:p["name"]}
	s.Init()
}

func main() {
}
