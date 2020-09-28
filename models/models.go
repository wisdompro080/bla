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
	Time string `json:"time,omitempty"`
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
