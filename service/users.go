package service

import (
	"errors"
	"fmt"
	"hotel/helpers"
	"hotel/model"
	"hotel/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *Service) Register(input RegisterInput) (UserInfo, error) {
	user := model.User{
		User_type: 2,
		Email:     input.Email,
		Password:  input.Password,
		Full_name: input.Full_name,
	}

	if err := utils.Validate.Struct(user); err != nil {
		return UserInfo{}, utils.NewError(utils.ErrFailedBind, err)
	}

	// create hashed password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserInfo{}, utils.NewError(utils.ErrInternalFailure, err)
	}
	user.Password = string(hashedPass)

	// register user
	result := s.DB.Omit("User_id").Create(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) { // broken, need pgconn
			return UserInfo{}, utils.NewError(utils.ErrBadRequest, err)
		}
		return UserInfo{}, utils.NewError(utils.ErrInternalFailure, err)
	}

	return UserInfo{
		Email:          input.Email,
		Password:       input.Password,
		Full_name:      input.Full_name,
		Deposit_amount: 0,
	}, nil
}

func (s *Service) Login(input LoginInput) (string, error) {
	// validate user input
	if err := utils.Validate.Struct(input); err != nil {
		return "", utils.NewError(utils.ErrFailedBind, err)
	}

	// check if user exists
	var existingUser model.User

	s.DB.Where("email = ?", input.Email).First(&existingUser)
	if existingUser.User_id == 0 {
		return "", utils.NewError(utils.ErrUnauthorized, fmt.Errorf("invalid email or password"))
	}

	// check password
	passCheck := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(input.Password))
	if passCheck != nil {
		return "", utils.NewError(utils.ErrUnauthorized, fmt.Errorf("invalid email or password"))
	}

	// generate JWT
	token, err := helpers.GenerateJWT(existingUser.User_id, existingUser.User_type)
	if err != nil {
		return "", utils.NewError(utils.ErrInternalFailure, err)
	}

	return token, nil
}
