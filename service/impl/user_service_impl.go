package impl

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"golang.org/x/crypto/bcrypt"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	repository.UserRepository
}

func (userService *userServiceImpl) Authentication(ctx context.Context, model model.UserModel) entity.User {
 	userResult, err := userService.UserRepository.Authentication(ctx, model.Username)
 	if err != nil {
 		panic(exception.UnauthorizedError{
 			Message: err.Error(),
 		})
 	}
 	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(model.Password))
 	if err != nil {
 		panic(exception.UnauthorizedError{
 			Message: "incorrect username and password",
 		})
 	}
 	return userResult
 }

func (userService *userServiceImpl) Register(ctx context.Context, model model.UserRegistrationModel) entity.User {
	// Check if email already exists
	_, err := userService.UserRepository.FindByEmail(ctx, model.Email)
	if err == nil {
		panic(exception.ValidationError{
			Message: "email already exists",
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)

	// Create user entity
	user := entity.User{
		Password:  string(hashedPassword),
		Name:      model.Name,
		Email:     model.Email,
		Role:      model.Role,
		IsActive:  true,
	}

	// Save user
	result, err := userService.UserRepository.Register(ctx, user)
	exception.PanicLogging(err)

	return result
}
