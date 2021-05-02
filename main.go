package main

import (
	"log"
	"todo-app/controller"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.POST("/add", controller.AddTodo)
	//e.POST("/getUserProfile", getUserProfile)
	//	e.GET("/getAllUsers", getAllUsers)
	//	e.PUT("/updateProfile", updateProfile)
	//	e.DELETE("/deleteProfile:id", deleteProfile)

	log.Fatal(e.Start(":8080"))

}
