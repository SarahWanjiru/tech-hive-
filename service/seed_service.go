package service

import (
	"context"
)

type SeedService interface {
	SeedUsers(ctx context.Context) error
	SeedProducts(ctx context.Context) error
	SeedAll(ctx context.Context) error
}