package config

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"test/models"
)

var Config models.Config
var Connection models.DatabaseConnection
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
	Connection.Col,Connection.Db=DbConnection()
}
func DbConnection() (driver.Collection, driver.Database) {

	var server string = Config.Arango.Server
	var database string = Config.Arango.Database
	var userName string = Config.Arango.UserName
	var password string = Config.Arango.Password
	var user string = Config.Arango.Collections.User
	log.Info(server)
	log.Info(userName)
	log.Info(password)

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{server},
	})
	if err != nil {
		log.Fatal("connection failed", err)
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(userName, password),
	})
	ctx := context.Background()
	db, err := client.Database(ctx, database)
	if err != nil {
		log.Fatal("db connetion error", err)
	}
	ctx = context.Background()
	col, err := db.Collection(ctx, user)
	if err != nil {
		log.Fatal("Collection creation Error", err)
	}
	log.Println("Connection Created")
	return col, db
}

