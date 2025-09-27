package impl

import (
 	"context"
 	"errors"
 	"github.com/tech-hive/ecommerce/entity"
 	"github.com/tech-hive/ecommerce/repository"
 	"gorm.io/gorm"
 )

func NewCartRepositoryImpl(DB *gorm.DB) repository.CartRepository {
	return &cartRepositoryImpl{DB: DB}
}

type cartRepositoryImpl struct {
	*gorm.DB
}

func (cartRepository *cartRepositoryImpl) GetCartByUserId(ctx context.Context, userId uint) (entity.Cart, error) {
	var cart entity.Cart
	result := cartRepository.DB.WithContext(ctx).
		Preload("CartItems").
		Preload("CartItems.Product").
		Where("user_id = ?", userId).
		First(&cart)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return entity.Cart{}, errors.New("cart not found")
		}
		return entity.Cart{}, result.Error
	}
	return cart, nil
}

func (cartRepository *cartRepositoryImpl) CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error) {
	result := cartRepository.DB.WithContext(ctx).Create(&cart)
	if result.Error != nil {
		return entity.Cart{}, result.Error
	}
	return cart, nil
}

func (cartRepository *cartRepositoryImpl) GetOrCreateCart(ctx context.Context, userId uint) (entity.Cart, error) {
	// Try to find existing cart
	cart, err := cartRepository.GetCartByUserId(ctx, userId)
	if err == nil {
		return cart, nil
	}

	// Create new cart if not found
	newCart := entity.Cart{
		UserId: userId,
	}
	return cartRepository.CreateCart(ctx, newCart)
}

func (cartRepository *cartRepositoryImpl) AddItemToCart(ctx context.Context, cartItem entity.CartItem) (entity.CartItem, error) {
	result := cartRepository.DB.WithContext(ctx).Create(&cartItem)
	if result.Error != nil {
		return entity.CartItem{}, result.Error
	}
	return cartItem, nil
}

func (cartRepository *cartRepositoryImpl) UpdateCartItem(ctx context.Context, cartItemId uint, quantity int32) (entity.CartItem, error) {
	var cartItem entity.CartItem
	result := cartRepository.DB.WithContext(ctx).Model(&cartItem).Where("id = ?", cartItemId).Update("quantity", quantity)
	if result.Error != nil {
		return entity.CartItem{}, result.Error
	}
	return cartItem, nil
}

func (cartRepository *cartRepositoryImpl) RemoveCartItem(ctx context.Context, cartItemId uint) error {
	result := cartRepository.DB.WithContext(ctx).Delete(&entity.CartItem{}, cartItemId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cartRepository *cartRepositoryImpl) ClearCart(ctx context.Context, cartId uint) error {
	result := cartRepository.DB.WithContext(ctx).Where("cart_id = ?", cartId).Delete(&entity.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}