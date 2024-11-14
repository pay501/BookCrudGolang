package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

type UserRepository interface {
	CreateUser(user *User) error
	LoginUser(user *User) (*User, error)
}
