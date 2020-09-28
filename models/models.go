package models

import (
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port   string
	Arango struct {
		Database    string
		Server      string
		UserName    string
		Password    string
		Collections struct {
			User string
		}
	}
	LogLevel log.Level
}
type Document struct {
	Key  string `json:"_key"`
	Time string `json:"time" binding:"required"`
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`

}
