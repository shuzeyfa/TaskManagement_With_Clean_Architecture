package router

import (
	"log"
	"taskmanagement/Delivery/controllers"
	infrastructure "taskmanagement/Infrastructure"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	infrastructure.ConnectDB()
	log.Println("App is ready!")

	r := gin.Default()

	r.POST("/register", controllers.RegisterHandler)
	r.POST("/login", controllers.LoginUser)

	r.GET("/tasks", infrastructure.AuthMiddleware(), controllers.GetAllTask)
	r.GET("/task/:id", infrastructure.AuthMiddleware(), controllers.GetTaskByID)
	r.POST("/create", infrastructure.AuthMiddleware(), controllers.CreateTask)
	r.PUT("/update/:id", infrastructure.AuthMiddleware(), controllers.UpdateTask)
	r.DELETE("/delete/:id", infrastructure.AuthMiddleware(), controllers.DeleteTask)

	return r
}
