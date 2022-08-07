package main

import (
	"net/http"
	"users-api-with-orm-with-jwt-auth/controllers"
	"users-api-with-orm-with-jwt-auth/middlewares"
	"users-api-with-orm-with-jwt-auth/models"
	conn "users-api-with-orm-with-jwt-auth/repository"

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

	r.POST("/login", controllers.Login)

	// s1 := r.Group("/v1/secure")
	// s1.Use(func(c *gin.Context) {
	// 	fmt.Print("executing middleware")
	// 	middlewares.IsAuthorized(c)
	// 	fmt.Print("executed middleware")

	// })

	// s1.GET("", func(ctx *gin.Context) {
	// 	fmt.Print("logged")

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"status": "welcome",
	// 	})
	// })

	r.GET("/secure",
		middlewares.IsAuthorized(),
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "welcome",
			})
		})

	r.Run()
}
