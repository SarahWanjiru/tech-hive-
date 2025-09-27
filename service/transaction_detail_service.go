package service

import (
	"context"
	"github.com/tech-hive/ecommerce/model"
)

type TransactionDetailService interface {
	FindById(ctx context.Context, id string) model.TransactionDetailModel
}
