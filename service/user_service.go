package service

import (
	"context"
	"github.com/tech-hive/ecommerce/entity"
	"github.com/tech-hive/ecommerce/model"
)

type UserService interface {
	Authentication(ctx context.Context, model model.UserModel) entity.User
	Register(ctx context.Context, model model.UserRegistrationModel) entity.User
}
