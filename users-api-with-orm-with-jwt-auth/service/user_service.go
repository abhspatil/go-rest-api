package service

import (
	"fmt"
	auth "users-api-with-orm-with-jwt-auth/auth"
	"users-api-with-orm-with-jwt-auth/dto"
	"users-api-with-orm-with-jwt-auth/models"
)

func (h handler) CreateUser(createUserReq *dto.CreateUserReq) *models.User {

	user := models.User{FullName: createUserReq.FullName, DOB: createUserReq.DOB,
		Salary: createUserReq.Salary, Password: createUserReq.Password,
		Role: createUserReq.Role, Email: createUserReq.Email,
	}
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

func (h handler) GetUserByEmail(email string) *models.User {

	var user models.User

	result := h.DB.Where("email = ?", email).First(&user)

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

	if updateUserReq.Password != "" {
		user.Password = updateUserReq.Password
	}

	if updateUserReq.Role != "" {
		user.Role = updateUserReq.Role
	}

	if updateUserReq.Email != "" {
		user.Email = updateUserReq.Email
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

func (h handler) SignIn(loinReq *dto.LoginReq) *dto.LoginResp {
	// var user *models.User
	var user *models.User = h.GetUserByEmail(loinReq.Email)

	if user.Id == 0 {
		fmt.Printf("user with email not found")
	}

	if user.Password != loinReq.Password {
		fmt.Printf("invalid email/password, please try again")
		return nil
	}

	validToken, err := auth.GenerateJWTToken(loinReq.Email, user.Role)
	if err != nil {
		// var err Error
		// err = SetError(err, "Failed to generate token")
		return nil
	}

	var loginResp dto.LoginResp
	loginResp.Id = user.Id
	loginResp.Email = user.Email
	loginResp.Role = user.Role
	loginResp.TokenString = validToken

	return &loginResp
}
