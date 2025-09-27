package repository

import (
	"context"
	"github.com/tech-hive/ecommerce/entity"
)

type TransactionDetailRepository interface {
	FindById(ctx context.Context, id string) (entity.TransactionDetail, error)
}
