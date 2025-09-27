package impl

import (
	"context"
	"errors"
	"strconv"
	"github.com/tech-hive/ecommerce/entity"
	"github.com/tech-hive/ecommerce/exception"
	"github.com/tech-hive/ecommerce/model"
	"github.com/tech-hive/ecommerce/repository"
	"github.com/tech-hive/ecommerce/service"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewOrderServiceImpl(orderRepository *repository.OrderRepository, cartRepository *repository.CartRepository, productRepository *repository.ProductRepository, DB *gorm.DB) service.OrderService {
	return &orderServiceImpl{
		OrderRepository:   *orderRepository,
		CartRepository:    *cartRepository,
		ProductRepository: *productRepository,
		DB:                DB,
	}
}

type orderServiceImpl struct {
	repository.OrderRepository
	repository.CartRepository
	repository.ProductRepository
	DB *gorm.DB
}

func (orderService *orderServiceImpl) CreateOrder(ctx context.Context, userId uint, request model.CreateOrderModel) (model.OrderModel, error) {
	// Get cart
	cart, err := orderService.CartRepository.GetCartByUserId(ctx, userId)
	if err != nil {
		return model.OrderModel{}, errors.New("cart not found")
	}

	if len(cart.CartItems) == 0 {
		return model.OrderModel{}, errors.New("cart is empty")
	}

	// Calculate total
	var total float64 = 0
	for _, item := range cart.CartItems {
		total += item.Price * float64(item.Quantity)
	}

	// Create order
	order := entity.Order{
		UserId: userId,
		Total:  total,
		Status: "pending",
	}

	// Use transaction to ensure data consistency
	tx := orderService.DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create order
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return model.OrderModel{}, err
	}

	// Create order items and update product stock
	var orderItems []entity.OrderItem
	for _, cartItem := range cart.CartItems {
		// Get product to check current stock
		product, err := orderService.ProductRepository.FindByProductId(ctx, cartItem.ProductId)
		if err != nil {
			tx.Rollback()
			return model.OrderModel{}, errors.New("product not found: " + cartItem.ProductId)
		}

		// Check stock availability
		if product.Stock < cartItem.Quantity {
			tx.Rollback()
			return model.OrderModel{}, errors.New("insufficient stock for product: " + product.Name)
		}

		// Create order item
		productUUID, err := uuid.Parse(cartItem.ProductId)
		if err != nil {
			tx.Rollback()
			return model.OrderModel{}, errors.New("invalid product ID format")
		}
		orderItem := entity.OrderItem{
			OrderId:   order.Id,
			ProductId: productUUID,
			Quantity:  cartItem.Quantity,
			Price:     cartItem.Price,
		}
		orderItems = append(orderItems, orderItem)

		// Update product stock
		if err := tx.Model(&product).Update("stock", gorm.Expr("stock - ?", cartItem.Quantity)).Error; err != nil {
			tx.Rollback()
			return model.OrderModel{}, err
		}
	}

	// Create order items
	if err := tx.Create(&orderItems).Error; err != nil {
		tx.Rollback()
		return model.OrderModel{}, err
	}

	// Clear cart
	if err := orderService.CartRepository.ClearCart(ctx, cart.Id); err != nil {
		tx.Rollback()
		return model.OrderModel{}, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return model.OrderModel{}, err
	}

	// Get created order with all details
	return orderService.GetOrderById(ctx, order.Id, userId)
}

func (orderService *orderServiceImpl) GetOrderById(ctx context.Context, orderId uint, userId uint) (model.OrderModel, error) {
	order, err := orderService.OrderRepository.GetOrderById(ctx, orderId)
	if err != nil {
		return model.OrderModel{}, err
	}

	// Check if order belongs to user
	if order.UserId != userId {
		return model.OrderModel{}, errors.New("order not found")
	}

	// Convert to model
	var orderItems []model.OrderItemModel
	for _, item := range order.OrderItems {
		orderItemModel := model.OrderItemModel{
			Id:        item.Id,
			OrderId:   item.OrderId,
			ProductId: item.ProductId.String(),
			Quantity:  item.Quantity,
			Price:     item.Price,
			CreatedAt: item.CreatedAt.String(),
			Product: model.ProductModel{
				Id:          strconv.FormatUint(uint64(item.Product.Id), 10),
				Name:        item.Product.Name,
				Description: item.Product.Description,
				Price:       item.Product.Price,
				Stock:       item.Product.Stock,
				ImageUrl:    item.Product.ImageUrl,
			},
		}
		orderItems = append(orderItems, orderItemModel)
	}

	// Convert payments
	var paymentModel *model.PaymentModel
	if len(order.Payments) > 0 {
		payment := order.Payments[0] // Assuming one payment per order
		paymentModel = &model.PaymentModel{
			Id:            payment.Id,
			OrderId:       payment.OrderId,
			TransactionId: payment.TransactionId,
			Status:        payment.Status,
			PaidAt:        payment.PaidAt.String(),
		}
	}

	orderModel := model.OrderModel{
		Id:         order.Id,
		UserId:     order.UserId,
		Total:      order.Total,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt.String(),
		OrderItems: orderItems,
		Payment:    paymentModel,
	}

	return orderModel, nil
}

func (orderService *orderServiceImpl) GetOrdersByUserId(ctx context.Context, userId uint) ([]model.OrderModel, error) {
	orders, err := orderService.OrderRepository.GetOrdersByUserId(ctx, userId)
	if err != nil {
		return []model.OrderModel{}, err
	}

	var orderModels []model.OrderModel
	for _, order := range orders {
		// Convert to model
		var orderItems []model.OrderItemModel
		for _, item := range order.OrderItems {
			orderItemModel := model.OrderItemModel{
				Id:        item.Id,
				OrderId:   item.OrderId,
				ProductId: item.ProductId.String(),
				Quantity:  item.Quantity,
				Price:     item.Price,
				CreatedAt: item.CreatedAt.String(),
				Product: model.ProductModel{
					Id:          strconv.FormatUint(uint64(item.Product.Id), 10),
					Name:        item.Product.Name,
					Description: item.Product.Description,
					Price:       item.Product.Price,
					Stock:       item.Product.Stock,
					ImageUrl:    item.Product.ImageUrl,
				},
			}
			orderItems = append(orderItems, orderItemModel)
		}

		// Convert payments
		var paymentModel *model.PaymentModel
		if len(order.Payments) > 0 {
			payment := order.Payments[0]
			paymentModel = &model.PaymentModel{
				Id:            payment.Id,
				OrderId:       payment.OrderId,
				TransactionId: payment.TransactionId,
				Status:        payment.Status,
				PaidAt:        payment.PaidAt.String(),
			}
		}

		orderModel := model.OrderModel{
			Id:         order.Id,
			UserId:     order.UserId,
			Total:      order.Total,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt.String(),
			OrderItems: orderItems,
			Payment:    paymentModel,
		}
		orderModels = append(orderModels, orderModel)
	}

	return orderModels, nil
}

func (orderService *orderServiceImpl) UpdateOrderStatus(ctx context.Context, orderId uint, request model.UpdateOrderStatusModel) (model.OrderModel, error) {
	// Get order first to check ownership
	_, err := orderService.OrderRepository.GetOrderById(ctx, orderId)
	if err != nil {
		return model.OrderModel{}, err
	}

	// Update order status
	updatedOrder, err := orderService.OrderRepository.UpdateOrderStatus(ctx, orderId, request.Status)
	if err != nil {
		return model.OrderModel{}, err
	}

	// Convert to model
	var orderItems []model.OrderItemModel
	for _, item := range updatedOrder.OrderItems {
		orderItemModel := model.OrderItemModel{
			Id:        item.Id,
			OrderId:   item.OrderId,
			ProductId: item.ProductId.String(),
			Quantity:  item.Quantity,
			Price:     item.Price,
			CreatedAt: item.CreatedAt.String(),
			Product: model.ProductModel{
				Id:          strconv.FormatUint(uint64(item.Product.Id), 10),
				Name:        item.Product.Name,
				Description: item.Product.Description,
				Price:       item.Product.Price,
				Stock:       item.Product.Stock,
				ImageUrl:    item.Product.ImageUrl,
			},
		}
		orderItems = append(orderItems, orderItemModel)
	}

	orderModel := model.OrderModel{
		Id:         updatedOrder.Id,
		UserId:     updatedOrder.UserId,
		Total:      updatedOrder.Total,
		Status:     updatedOrder.Status,
		CreatedAt:  updatedOrder.CreatedAt.String(),
		OrderItems: orderItems,
	}

	return orderModel, nil
}

func (orderService *orderServiceImpl) CancelOrder(ctx context.Context, orderId uint, userId uint) error {
	// Get order first to check ownership
	order, err := orderService.OrderRepository.GetOrderById(ctx, orderId)
	if err != nil {
		return err
	}

	// Check if order belongs to user
	if order.UserId != userId {
		return errors.New("order not found")
	}

	// Check if order can be cancelled
	if order.Status == "shipped" || order.Status == "delivered" {
		return errors.New("cannot cancel shipped or delivered order")
	}

	// Update order status to cancelled
	_, err = orderService.OrderRepository.UpdateOrderStatus(ctx, orderId, "cancelled")
	if err != nil {
		return err
	}

	// If order was paid, you might want to implement refund logic here
	// For now, we'll just cancel the order

	return nil
}