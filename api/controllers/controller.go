package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/TheMrPuffin/linkdex/configs"
	"github.com/TheMrPuffin/linkdex/models"
	"github.com/TheMrPuffin/linkdex/responses"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var linkCollection *mongo.Collection = configs.GetLinksCollection(configs.DB)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetAllLinks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var links []models.Link
		defer cancel()

		results, err := linkCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var link models.Link
			if err = results.Decode(&link); err != nil {
				c.JSON(http.StatusInternalServerError, responses.LinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			links = append(links, link)
		}

		c.JSON(http.StatusOK,
			responses.LinkResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": links}},
		)
	}
}
