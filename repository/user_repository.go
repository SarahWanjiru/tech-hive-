package repository

import (
	"context"
	"github.com/tech-hive/ecommerce/entity"
)

type UserRepository interface {
	Authentication(ctx context.Context, email string) (entity.User, error)
	Register(ctx context.Context, user entity.User) (entity.User, error)
	FindByUsername(ctx context.Context, username string) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	Create(email string, password string, roles []string)
	DeleteAll()
}
