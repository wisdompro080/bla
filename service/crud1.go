package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"test/models"

	//log "github.com/sirupsen/logrus"
	"test/config"
)

func Create1(c *gin.Context) {
	_, db := DbConnection()
	ctx := context.Background()
	var details models.Document
	err := c.BindJSON(&details)
	if err != nil {
		log.Fatal(err)
	}
	collectionName := config.Config.Arango.Collections.User
	query := "INSERT {name:" + "'" + details.Name + "'" + ",id:'" + details.Id + "'} IN " + collectionName
	fmt.Println(query)
	_, _ = db.Query(ctx, query, nil)
}
func Read1(c *gin.Context) {
	DbConnection()
	_, db := DbConnection()
	ctx := context.Background()
	query := "FOR d IN Documents RETURN d"
	cursor, err := db.Query(ctx, query, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close()
	for {
		var doc models.Document
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		jdoc, err := json.Marshal(doc)
		if err != nil {
			log.Fatal("marshal error ", err)
		}
		c.String(200, string(jdoc))
	}
}
func Remove1(c *gin.Context) {
	_, db := DbConnection()
	ctx := context.Background()
	collectionName := config.Config.Arango.Collections.User
	key := c.Param("id")
	query := "REMOVE'" + key + "'IN " + collectionName
	_, err := db.Query(ctx, query, nil)
	if err != nil {
		log.Fatal("error in removal", err)
	}
	log.Info("Removed Succesfully")
}
func Update1(c *gin.Context) {
	DbConnection()
	_, db := DbConnection()
	ctx := context.Background()
	var doc models.Document
	_ = c.ShouldBindWith(&doc, binding.JSON)
	key := c.Param("id")
	collectionName := config.Config.Arango.Collections.User
	query := "REPLACE'" + key + "'WITH{id:'" + doc.Id + "',name:'" + doc.Name + "'} IN " + collectionName
	_, err := db.Query(ctx, query, nil)
	if err != nil {
		log.Fatal("error in removal", err)
	}
	log.Info("updated successfully", doc)
}
