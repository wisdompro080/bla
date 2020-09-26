package service

import (
	"context"
	"encoding/json"
	"github.com/arangodb/go-driver"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"test/models"
)

func Create(c *gin.Context) {
	col, _ := DbConnection()
	ctx := context.Background()
	var p models.Document
	err := c.ShouldBindWith(&p, binding.JSON)
	if err != nil {
		log.Fatal("Binding not successfull", err)
	}
	_, err = col.CreateDocument(ctx, p)
	if err != nil {
		log.Fatal("Error in creating the given document", err)
	}
	log.Info("Created document", p)
}

func Read(c *gin.Context) {
	DbConnection()
	_, db := DbConnection()
	ctx := context.Background()
	query := "FOR d IN Documents RETURN d"
	cursor, err := db.Query(ctx, query, nil)

	if err != nil {
		// handle error
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

func Remove(c *gin.Context) {
	DbConnection()
	col, _ := DbConnection()
	key := c.Param("id")
	ctx := context.Background()
	//fmt.Println(key)
	_, err := col.RemoveDocument(ctx, key)
	if err != nil {
		c.String(500, "Error Enter a valid id")
		log.Panic("error in removing document", err)
	}
	log.Info("document removed successfully")
}

func Update(c *gin.Context) {
	DbConnection()
	col, _ := DbConnection()
	ctx := context.Background()
	//patch := make(map[string]interface{})
	//patch["name"] = "raul"
	//patch["Id"] = "234"
	var doc models.Document
	_ = c.ShouldBindWith(&doc, binding.JSON)
	key := c.Param("id")
	_, err := col.UpdateDocument(ctx, key, doc)
	if err != nil {
		log.Fatal("Unable to update", err)
	}
	log.Info("updated successfully", doc)
}
