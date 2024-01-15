package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func CreateUser(db *gorm.DB, user *User) error {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPwd)

	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func LoginUser(db *gorm.DB, user *User) (string, error) {
	selectedUser := new(User)

	result := db.Where("email = ?", user.Email).First(selectedUser)
	if result.Error != nil {
		fmt.Println("Error on database")
		return "", result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		fmt.Println("Compare password")
		return "", err
	}

	// Create JWT token
	jwtSecretKey := "testSecret"
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectedUser.ID
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		fmt.Println("JWT")
		return "", err
	}

	return t, nil
}
