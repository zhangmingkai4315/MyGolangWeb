package utils

import (
	"github.com/spf13/viper"
	"log"
)
var Env string
func InitConfigFile() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("Loading config file error.", err)
	} else {
		log.Println("Load config file success")
	}
	Env=viper.GetString("env")
	log.Println("Setting server running env to",Env)
}
