package impl

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/repository"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (userRepository *userRepositoryImpl) Create(email string, password string, roles []string) {
	// This method is deprecated - use Register instead
	// For backward compatibility, we'll create a user with customer role
	user := entity.User{
		Password: password,
		Name:     "User", // Default name for backward compatibility
		Email:    email,
		Role:     "customer",
		IsActive: true,
	}
	err := userRepository.DB.Create(&user).Error
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) DeleteAll() {
	err := userRepository.DB.Where("1=1").Delete(&entity.User{}).Error
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) Authentication(ctx context.Context, email string) (entity.User, error) {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).
		Where("email = ? and is_active = ?", email, true).
		First(&userResult)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, result.Error
	}
	return userResult, nil
}

func (userRepository *userRepositoryImpl) Register(ctx context.Context, user entity.User) (entity.User, error) {
	result := userRepository.DB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (userRepository *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	result := userRepository.DB.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (userRepository *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	result := userRepository.DB.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}
