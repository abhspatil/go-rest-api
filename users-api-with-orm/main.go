package main

import (
	"net/http"
	"users-api-with-orm/controllers"
	"users-api-with-orm/models"
	conn "users-api-with-orm/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	DB := conn.SetupDB()
	DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", DB)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "ok"})
	})

	r.GET("/users", controllers.GetAllUsers)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users", controllers.UpdateUser)
	r.GET("/users/:userId", controllers.GetUserById)
	r.DELETE("/users/:userId", controllers.DeleteById)
	r.Run()
}
