package repository

import (
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryDB{db: db}
}

func (repo *userRepositoryDB) CreateUser(user *User) error {

	result := repo.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *userRepositoryDB) LoginUser(user *User) (*User, error) {
	selectedUser := User{}

	result := repo.db.Where("email = ?", user.Email).First(&selectedUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &selectedUser, nil
}
