package service

import (
	"fmt"
	"users-api-with-orm/dto"
	"users-api-with-orm/models"
)

func (h handler) CreateUser(createUserReq *dto.CreateUserReq) *models.User {

	user := models.User{FullName: createUserReq.FullName, DOB: createUserReq.DOB, Salary: createUserReq.Salary}
	h.DB.Create(&user)

	// db := c.MustGet("db").(*gorm.DB)
	// db.Create(&user)
	return &user
}

func (h handler) GetUserById(userId int) *models.User {

	var user models.User

	result := h.DB.First(&user, userId)

	if result.Error != nil {
		fmt.Println(result.Error)
		panic(result.Error)
	}

	return &user
}

func (h handler) DeleteById(userId int) {

	var user models.User

	result := h.DB.First(&user, userId)

	if result.Error != nil {
		fmt.Println(result.Error)
		panic(result.Error)
	}

	h.DB.Delete(&user)
}

func (h handler) Update(updateUserReq *dto.UpdateUserReq) *models.User {

	var user models.User

	result := h.DB.First(&user, updateUserReq.Id)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println("fromDB::", user)
	fmt.Println("fromReq::", updateUserReq)

	if updateUserReq.FullName != "" {
		user.FullName = updateUserReq.FullName
	}

	if updateUserReq.DOB.IsZero() {
		user.DOB = updateUserReq.DOB
	}

	if updateUserReq.Salary != 0.0 {
		user.Salary = updateUserReq.Salary
	}

	h.DB.Save(&user)

	return &user
}

func (h handler) GetAll() []models.User {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		fmt.Println(result.Error)
	}

	return users
}
