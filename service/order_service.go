package service

import (
	"context"
	"github.com/tech-hive/ecommerce/model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userId uint, request model.CreateOrderModel) (model.OrderModel, error)
	GetOrderById(ctx context.Context, orderId uint, userId uint) (model.OrderModel, error)
	GetOrdersByUserId(ctx context.Context, userId uint) ([]model.OrderModel, error)
	UpdateOrderStatus(ctx context.Context, orderId uint, request model.UpdateOrderStatusModel) (model.OrderModel, error)
	CancelOrder(ctx context.Context, orderId uint, userId uint) error
}