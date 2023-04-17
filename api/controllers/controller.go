package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/TheMrPuffin/linkdex/configs"
	"github.com/TheMrPuffin/linkdex/models"
	"github.com/TheMrPuffin/linkdex/responses"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var linkCollection *mongo.Collection = configs.GetLinksCollection(configs.DB)
var validate = validator.New()

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
			responses.LinkResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"links": links}},
		)
	}
}

func GetLinkByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		linkName := c.Param("linkName")
		var link models.Link
		defer cancel()

		err := linkCollection.FindOne(ctx, bson.M{"name": linkName}).Decode(&link)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.LinkResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"link": link}})
	}
}

func CreateLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var link models.Link
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&link); err != nil {
			c.JSON(http.StatusBadRequest, responses.LinkResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&link); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LinkResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newLink := models.Link{
			Name: link.Name,
			Url:  link.Url,
		}

		result, err := linkCollection.InsertOne(ctx, newLink)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.LinkResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DeleteLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		linkName := c.Param("linkName")
		defer cancel()

		result, err := linkCollection.DeleteOne(ctx, bson.M{"name": linkName})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.LinkResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Link with specified name not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.LinkResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Link successfully deleted!"}},
		)
	}
}
