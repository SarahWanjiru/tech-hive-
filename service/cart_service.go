package service

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
)

type CartService interface {
	GetCart(ctx context.Context, userId uint) (model.CartModel, error)
	AddToCart(ctx context.Context, userId uint, request model.AddToCartModel) (model.CartItemModel, error)
	UpdateCartItem(ctx context.Context, userId uint, cartItemId uint, request model.UpdateCartItemModel) (model.CartItemModel, error)
	RemoveFromCart(ctx context.Context, userId uint, cartItemId uint) error
	ClearCart(ctx context.Context, userId uint) error
}