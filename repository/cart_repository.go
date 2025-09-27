package repository

import (
	"context"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/entity"
)

type CartRepository interface {
	GetCartByUserId(ctx context.Context, userId uint) (entity.Cart, error)
	CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error)
	GetOrCreateCart(ctx context.Context, userId uint) (entity.Cart, error)
	AddItemToCart(ctx context.Context, cartItem entity.CartItem) (entity.CartItem, error)
	UpdateCartItem(ctx context.Context, cartItemId uint, quantity int32) (entity.CartItem, error)
	RemoveCartItem(ctx context.Context, cartItemId uint) error
	ClearCart(ctx context.Context, cartId uint) error
}