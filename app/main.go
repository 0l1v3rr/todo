package main

import (
	"fmt"

	"github.com/0l1v3rr/todo/app/controller"
	"github.com/0l1v3rr/todo/app/data"
	"github.com/0l1v3rr/todo/app/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// loading the environment variables
	err := data.GetVariables()
	if err != nil {
		fmt.Println("An error occurred while reading the .env file: ")
		fmt.Println(err.Error())
		return
	}

	// connecting to the db
	err = model.Setup()
	if err != nil {
		fmt.Println("Failed to connect to the database: ")
		fmt.Println(err.Error())
		return
	}

	// creating the gin router
	r := gin.Default()

	// using the cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// specifying the login system endpoints
	r.POST("/api/v1/register", controller.Register)
	r.POST("/api/v1/login", controller.Login)
	r.POST("/api/v1/logout")

	// running the router
	r.Run(fmt.Sprintf(":%s", data.Env["PORT"]))
}
