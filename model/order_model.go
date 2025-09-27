package model

type OrderModel struct {
	Id         uint                   `json:"id"`
	UserId     uint                   `json:"user_id"`
	Total      float64                `json:"total"`
	Status     string                 `json:"status"`
	CreatedAt  string                 `json:"created_at"`
	OrderItems []OrderItemModel       `json:"order_items"`
	Payment    *PaymentModel          `json:"payment,omitempty"`
}

type OrderItemModel struct {
	Id         uint    `json:"id"`
	OrderId    uint    `json:"order_id"`
	ProductId  string  `json:"product_id"`
	Product    ProductModel `json:"product"`
	Quantity   int32   `json:"quantity"`
	Price      float64 `json:"price"`
	CreatedAt  string  `json:"created_at"`
}

type CreateOrderModel struct {
	CartId     uint    `json:"cart_id" validate:"required"`
	ShippingAddress string `json:"shipping_address" validate:"required"`
}

type UpdateOrderStatusModel struct {
	Status string `json:"status" validate:"required,oneof=pending confirmed processing shipped delivered cancelled"`
}

type PaymentModel struct {
	Id            uint   `json:"id"`
	OrderId       uint   `json:"order_id"`
	TransactionId string `json:"transaction_id"`
	Status        string `json:"status"`
	PaidAt        string `json:"paid_at"`
}