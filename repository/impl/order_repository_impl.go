package impl

import (
 	"context"
 	"errors"
 	"github.com/tech-hive/ecommerce/entity"
 	"github.com/tech-hive/ecommerce/repository"
 	"gorm.io/gorm"
 )

func NewOrderRepositoryImpl(DB *gorm.DB) repository.OrderRepository {
	return &orderRepositoryImpl{DB: DB}
}

type orderRepositoryImpl struct {
	*gorm.DB
}

func (orderRepository *orderRepositoryImpl) CreateOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	result := orderRepository.DB.WithContext(ctx).Create(&order)
	if result.Error != nil {
		return entity.Order{}, result.Error
	}
	return order, nil
}

func (orderRepository *orderRepositoryImpl) GetOrderById(ctx context.Context, orderId uint) (entity.Order, error) {
	var order entity.Order
	result := orderRepository.DB.WithContext(ctx).
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("Payments").
		Where("id = ?", orderId).
		First(&order)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return entity.Order{}, errors.New("order not found")
		}
		return entity.Order{}, result.Error
	}
	return order, nil
}

func (orderRepository *orderRepositoryImpl) GetOrdersByUserId(ctx context.Context, userId uint) ([]entity.Order, error) {
	var orders []entity.Order
	result := orderRepository.DB.WithContext(ctx).
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Preload("Payments").
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Find(&orders)

	if result.Error != nil {
		return []entity.Order{}, result.Error
	}
	return orders, nil
}

func (orderRepository *orderRepositoryImpl) UpdateOrderStatus(ctx context.Context, orderId uint, status string) (entity.Order, error) {
	var order entity.Order
	result := orderRepository.DB.WithContext(ctx).Model(&order).Where("id = ?", orderId).Update("status", status)
	if result.Error != nil {
		return entity.Order{}, result.Error
	}

	// Return updated order
	return orderRepository.GetOrderById(ctx, orderId)
}

func (orderRepository *orderRepositoryImpl) DeleteOrder(ctx context.Context, orderId uint) error {
	result := orderRepository.DB.WithContext(ctx).Delete(&entity.Order{}, orderId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}