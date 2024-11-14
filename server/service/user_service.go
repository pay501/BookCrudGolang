package model

import (
	"fmt"
	"spy/errs"
	"spy/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) SignUp(data *repository.User) error {

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	data.Password = string(hashPwd)

	err = s.userRepo.CreateUser(data)
	if err != nil {
		return errs.NewUnexpected("unexpected error")
	}

	return nil
}

func (s *userService) Login(data *LoginReq) (string, error) {
	//selectedUser := new(User)
	user := repository.User{
		Email:    data.Email,
		Password: data.Password,
	}
	resultUser, err := s.userRepo.LoginUser(&user)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(data.Password))
	if err != nil {
		fmt.Println("Compare password")
		return "", err
	}

	// Create JWT token
	jwtSecretKey := "testSecret"
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = resultUser.ID
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		fmt.Println("JWT")
		return "", err
	}

	return t, nil
}
