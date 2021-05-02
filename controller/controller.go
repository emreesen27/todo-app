package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"todo-app/configuration"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Todo struct {
	Title     string `json:"title"`
	Completed *bool  `json:"completed"`
}

func AddTodo(c echo.Context) error {

	config := new(configuration.MongoConfiguration).Init(getDBUri(), getDBName())
	collection := config.Database().Collection("todos")

	todo := Todo{}
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, &todo)

	if todo.Completed == nil || todo.Title == "" {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	_, err = collection.InsertOne(context.TODO(), bson.D{{Key: todo.Title, Value: todo.Completed}})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success")
}

func getDBUri() string {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return "mongodb://localhost:27017"
	}
	return uri
}

func getDBName() string {
	return "todo-service"
}
