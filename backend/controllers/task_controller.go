package controllers

import (
	"context"
	"net/http"

	"github.com/ericengeng/tasks-app/backend/config"
	"github.com/ericengeng/tasks-app/backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTasks(c *gin.Context) {
	// build a list of tasks
	var tasks []models.Task
	//call get collection to retrieve the collection
	collection := config.GetCollection(config.ConnectDB(), "tasks")

	//get everything in the collection.  bson.M is all, with default context
	cursor, err := collection.Find(context.Background(), bson.M{})
	//if we got an error, run a status
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//cleans up and releases resources after done.  Put it here in case we
	//return after the error?
	defer cursor.Close(context.Background())

	//iterate
	for cursor.Next(context.Background()) {
		//decode and add to our output
		var task models.Task
		cursor.Decode((&task))
		tasks = append(tasks, task)
	}
	//
	c.JSON(http.StatusOK, tasks)

}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := config.GetCollection(config.ConnectDB(), "tasks")
	result, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"insertedID": result.InsertedID})
}
