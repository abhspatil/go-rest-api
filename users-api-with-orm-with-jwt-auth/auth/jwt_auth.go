package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("supersecretkey@Patil")

func GenerateJWTToken(email, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute).Unix()

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		// fmt.Errorf("Something Went Wrong: %s", err.Error())
		fmt.Printf("Something went wrong in creating jwt: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token, %q", token.Header["alg"])

		}
		return jwtKey, nil
	})
}
