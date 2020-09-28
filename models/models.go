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
type StudentDetails struct {
	Key        string `json:"_key" binding:"required"`
	Time       string `json:"time" binding:"required"`
	Id         string `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Rollnumber string `json:"rollnumber" binding:"required"`
	Marks      struct {
		Phy       string `json:"phy" binding:"required"`
		Chemistry string `json:"chem" binding:"required"`
		Maths     string `json:"maths" binding:"required"`
	} `json:"marks"`
}
