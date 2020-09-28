package service
//
//import (
//	"context"
//	"github.com/arangodb/go-driver"
//	"github.com/arangodb/go-driver/http"
//	log "github.com/sirupsen/logrus"
//	"test/config"
//)
//
//func DbConnection() (driver.Collection, driver.Database) {
//
//	var server string = config.Config.Arango.Server
//	var database string = config.Config.Arango.Database
//	var userName string = config.Config.Arango.UserName
//	var password string = config.Config.Arango.Password
//	var user string = config.Config.Arango.Collections.User
//	log.Info(server)
//	log.Info(userName)
//	log.Info(password)
//
//	conn, err := http.NewConnection(http.ConnectionConfig{
//		Endpoints: []string{server},
//	})
//	if err != nil {
//		log.Fatal("connection failed", err)
//	}
//	client, err := driver.NewClient(driver.ClientConfig{
//		Connection:     conn,
//		Authentication: driver.BasicAuthentication(userName, password),
//	})
//	ctx := context.Background()
//	db, err := client.Database(ctx, database)
//	if err != nil {
//		log.Fatal("db connetion error", err)
//	}
//	ctx = context.Background()
//	col, err := db.Collection(ctx, user)
//	if err != nil {
//		log.Fatal("Collection creation Error", err)
//	}
//	log.Println("Connection Created")
//	return col, db
//}
