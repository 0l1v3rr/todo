package main

import (
	"fmt"
	"os"

	"github.com/0l1v3rr/todo/app/controller"
	_ "github.com/0l1v3rr/todo/app/docs"
	"github.com/0l1v3rr/todo/app/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Advanced ToDo application
// @version         1.0
// @description     This is the API of the advanced ToDo application

// @contact.name	API Support
// @contact.url 	https://0l1v3rr.github.io
// @contact.email 	oliver.mrakovics@gmail.com

// @license.name 	MIT
// @license.url 	https://opensource.org/licenses/MIT

// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	// loading the environment variables
	godotenv.Load(".env")

	// connecting to the db
	err := model.Setup()
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
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// user endpoints
	r.GET("/api/v1/user", controller.GetLoggedInUser)
	r.POST("/api/v1/register", controller.Register)
	r.POST("/api/v1/login", controller.Login)
	r.POST("/api/v1/logout", controller.Logout)

	// task enpoints
	r.GET("/api/v1/tasks/list/:listId", controller.GetTasksByListId)
	r.GET("/api/v1/tasks/:url", controller.GetTaskByUrl)
	r.POST("/api/v1/tasks", controller.CreateTask)
	r.PATCH("/api/v1/tasks/:id", controller.ChangeTaskStatus)
	r.PUT("/api/v1/tasks/:id", controller.EditTask)
	r.DELETE("/api/v1/tasks/:id", controller.DeleteTask)

	// list endpoints
	r.GET("/api/v1/lists/user/:userId", controller.GetListsByUserId)
	r.GET("/api/v1/lists/:url", controller.GetListByUrl)
	r.POST("/api/v1/lists", controller.CreateList)

	// file endpoints
	r.POST("/api/v1/files", controller.UploadFile)

	// serving the static images
	r.Static("/assets/images", "./images")

	// swagger init
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// running the router
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
