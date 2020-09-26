package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"test/models"
)

var Config models.Config

func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.ReadInConfig()
	Config.Port = viper.GetString("PORT")
	log.Info(Config.Port)
	Config.Arango.Database = viper.GetString("DATABASE")
	Config.Arango.Server = viper.GetString("SERVER")
	Config.Arango.UserName = viper.GetString("USERNAME")
	Config.Arango.Password = viper.GetString("PASSWORD")
	Config.Arango.Collections.User = viper.GetString("USER")
	LogLevel := viper.GetString("LOGLEVEL")
	Level, err := log.ParseLevel(LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	Config.LogLevel = Level
}
