package controllers

import (
	"net/http"
	"strconv"
	"users-api-with-orm/dto"
	conn "users-api-with-orm/repository"
	"users-api-with-orm/service"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	DB := conn.SetupDB()
	h := service.New(DB)

	c.JSON(http.StatusOK, gin.H{"users": h.GetAll()})
}

func CreateUser(c *gin.Context) {
	DB := conn.SetupDB()
	h := service.New(DB)

	var createUserReq dto.CreateUserReq

	if err := c.ShouldBindJSON(&createUserReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": h.CreateUser(&createUserReq)})
}

func GetUserById(c *gin.Context) {
	DB := conn.SetupDB()
	h := service.New(DB)

	userId := c.Param("userId")

	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"user": h.GetUserById(userIdInt)})
}

func UpdateUser(c *gin.Context) {
	DB := conn.SetupDB()
	h := service.New(DB)

	var updateUserReq dto.UpdateUserReq

	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": h.Update(&updateUserReq)})
}

func DeleteById(c *gin.Context) {
	DB := conn.SetupDB()
	h := service.New(DB)

	userId := c.Param("userId")

	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		panic(err)
	}

	h.DeleteById(userIdInt)
	c.JSON(http.StatusOK, gin.H{})
}
