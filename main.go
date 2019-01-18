package main

import (
	"fmt"
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
	v := viper.GetStringMapString("project")
	projectPath = v["name"]
}

func main() {
}
