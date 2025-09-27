package model

type CartModel struct {
	Id        uint                   `json:"id"`
	UserId    uint                   `json:"user_id"`
	Items     []CartItemModel        `json:"items"`
	Total     float64                `json:"total"`
	CreatedAt string                 `json:"created_at"`
}

type CartItemModel struct {
	Id         uint    `json:"id"`
	CartId     uint    `json:"cart_id"`
	ProductId  string  `json:"product_id"`
	Product    ProductModel `json:"product"`
	Quantity   int32   `json:"quantity"`
	Price      float64 `json:"price"`
	CreatedAt  string  `json:"created_at"`
}

type AddToCartModel struct {
	ProductId string `json:"product_id" validate:"required"`
	Quantity  int32  `json:"quantity" validate:"required,min=1"`
}

type UpdateCartItemModel struct {
	Quantity int32 `json:"quantity" validate:"required,min=1"`
}