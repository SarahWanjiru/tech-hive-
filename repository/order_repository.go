package repository

import (
	"context"
	"github.com/tech-hive/ecommerce/entity"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order entity.Order) (entity.Order, error)
	GetOrderById(ctx context.Context, orderId uint) (entity.Order, error)
	GetOrdersByUserId(ctx context.Context, userId uint) ([]entity.Order, error)
	UpdateOrderStatus(ctx context.Context, orderId uint, status string) (entity.Order, error)
	DeleteOrder(ctx context.Context, orderId uint) error
}