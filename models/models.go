package models

import (
	"github.com/arangodb/go-driver"
	log "github.com/sirupsen/logrus"
)

type DatabaseConnection struct {
	Db  driver.Database
	Col driver.Collection
}
type Config struct {
	Port   string `json:"port"`
	Arango struct {
		Database    string `json:"database"`
		Server      string `json:"server"`
		UserName    string `json:"userName"`
		Password    string `json:"password"`
		Collections struct {
			User string `json:"user"`
		} `json:"collections"`
	} `json:"arango"`
	LogLevel log.Level `json:"logLevel"`
}
type StudentDetails struct {
	Key        string `json:"_key" binding:"required"`
	Time       string `json:"time" binding:"required"`
	Id         string `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	RollNumber string `json:"rollNumber" binding:"required"`
	Marks      struct {
		Phy       string `json:"phy" binding:"required"`
		Chemistry string `json:"chem" binding:"required"`
		Maths     string `json:"maths" binding:"required"`
	} `json:"marks"`
}
