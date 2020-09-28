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
	"time"

	//log "github.com/sirupsen/logrus"
	"test/config"
)

func Create1(c *gin.Context) {
	_, db := DbConnection()
	ctx := context.Background()
	var details models.Document
	err := c.BindJSON(&details)

	t := time.Now()
	details.Time = t.Format(time.RFC3339)
	a, _ := json.Marshal(details)
	fmt.Println(string(a))

	if err != nil {
		log.Fatal(err)
	}
	collectionName := config.Config.Arango.Collections.User
	bytecode, err := json.Marshal(details)
	fmt.Println(string(bytecode))
	query := "INSERT" + string(bytecode) + "IN " + collectionName

	_, _ = db.Query(ctx, query, nil)
}

func Read1(c *gin.Context) {
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
func ReadId1(c *gin.Context) {
	_, db := DbConnection()
	collectionName := config.Config.Arango.Collections.User
	key := c.Param("id")
	fmt.Println(key)
	ctx := context.Background()
	query := "for i in " + collectionName + " filter i._key=='" + key + "' return i"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		log.Fatal("key not present")
	}
	var doc models.Document
	_, err = cursor.ReadDocument(ctx, &doc)
	bytecode, err := json.Marshal(doc)
	if err != nil {
		log.Fatal("key not present")
	}
	c.String(200, string(bytecode))
	cursor.Close()

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
	_, db := DbConnection()
	ctx := context.Background()
	var doc models.Document
	_ = c.ShouldBindWith(&doc, binding.JSON)
	key := c.Param("id")
	collectionName := config.Config.Arango.Collections.User
	t := time.Now()
	doc.Time = t.Format(time.RFC3339)
	byteData, _ := json.Marshal(doc)
	query := "for i in " + collectionName + " filter i._key=='" + key + "' UPDATE i with" + string(byteData) + "IN " + collectionName
	_, err := db.Query(ctx, query, nil)

	if err != nil {
		log.Fatal("error in removal", err)
	}
	log.Info("updated successfully", doc)
}
