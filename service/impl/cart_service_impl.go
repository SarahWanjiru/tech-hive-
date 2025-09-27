package impl

import (
	"context"
	"errors"
	"strconv"
	"github.com/tech-hive/ecommerce/entity"
	"github.com/tech-hive/ecommerce/model"
	"github.com/tech-hive/ecommerce/repository"
	"github.com/tech-hive/ecommerce/service"
	"gorm.io/gorm"
)

func NewCartServiceImpl(cartRepository *repository.CartRepository, productRepository *repository.ProductRepository, DB *gorm.DB) service.CartService {
	return &cartServiceImpl{
		CartRepository:    *cartRepository,
		ProductRepository: *productRepository,
		DB:                DB,
	}
}

type cartServiceImpl struct {
	repository.CartRepository
	repository.ProductRepository
	DB *gorm.DB
}

func (cartService *cartServiceImpl) GetCart(ctx context.Context, userId uint) (model.CartModel, error) {
	cart, err := cartService.CartRepository.GetCartByUserId(ctx, userId)
	if err != nil {
		if err.Error() == "cart not found" {
			// Return empty cart if not found
			return model.CartModel{
				UserId: userId,
				Items:  []model.CartItemModel{},
				Total:  0,
			}, nil
		}
		return model.CartModel{}, err
	}

	// Convert to model
	var cartItems []model.CartItemModel
	var total float64 = 0

	for _, item := range cart.CartItems {
		cartItemModel := model.CartItemModel{
			Id:        item.Id,
			CartId:    item.CartId,
			ProductId: item.ProductId,
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
		cartItems = append(cartItems, cartItemModel)
		total += item.Price * float64(item.Quantity)
	}

	cartModel := model.CartModel{
		Id:        cart.Id,
		UserId:    cart.UserId,
		Items:     cartItems,
		Total:     total,
		CreatedAt: cart.CreatedAt.String(),
	}

	return cartModel, nil
}

func (cartService *cartServiceImpl) AddToCart(ctx context.Context, userId uint, request model.AddToCartModel) (model.CartItemModel, error) {
	// Get or create cart
	cart, err := cartService.CartRepository.GetOrCreateCart(ctx, userId)
	if err != nil {
		return model.CartItemModel{}, err
	}

	// Get product
	product, err := cartService.ProductRepository.FindByProductId(ctx, request.ProductId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.CartItemModel{}, errors.New("product not found")
		}
		return model.CartItemModel{}, err
	}

	// Check stock availability
	if product.Stock < request.Quantity {
		return model.CartItemModel{}, errors.New("insufficient stock")
	}

	// Check if item already exists in cart
	var existingItem entity.CartItem
	result := cartService.DB.WithContext(ctx).
		Where("cart_id = ? AND product_id = ?", cart.Id, request.ProductId).
		First(&existingItem)

	if result.Error == nil {
		// Update existing item
		newQuantity := existingItem.Quantity + request.Quantity
		if product.Stock < newQuantity {
			return model.CartItemModel{}, errors.New("insufficient stock for updated quantity")
		}

		_, err := cartService.CartRepository.UpdateCartItem(ctx, existingItem.Id, newQuantity)
		if err != nil {
			return model.CartItemModel{}, err
		}

		// Return updated item
		return model.CartItemModel{
			Id:        existingItem.Id,
			CartId:    existingItem.CartId,
			ProductId: existingItem.ProductId,
			Quantity:  newQuantity,
			Price:     existingItem.Price,
			CreatedAt: existingItem.CreatedAt.String(),
			Product: model.ProductModel{
				Id:          strconv.FormatUint(uint64(product.Id), 10),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Stock:       product.Stock,
				ImageUrl:    product.ImageUrl,
			},
		}, nil
	}

	// Create new cart item
	cartItem := entity.CartItem{
		CartId:    cart.Id,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
		Price:     product.Price,
	}

	resultItem, err := cartService.CartRepository.AddItemToCart(ctx, cartItem)
	if err != nil {
		return model.CartItemModel{}, err
	}

	return model.CartItemModel{
		Id:        resultItem.Id,
		CartId:    resultItem.CartId,
		ProductId: resultItem.ProductId,
		Quantity:  resultItem.Quantity,
		Price:     resultItem.Price,
		CreatedAt: resultItem.CreatedAt.String(),
		Product: model.ProductModel{
			Id:          strconv.FormatUint(uint64(product.Id), 10),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			ImageUrl:    product.ImageUrl,
		},
	}, nil
}

func (cartService *cartServiceImpl) UpdateCartItem(ctx context.Context, userId uint, cartItemId uint, request model.UpdateCartItemModel) (model.CartItemModel, error) {
	// Get cart to verify ownership
	cart, err := cartService.CartRepository.GetCartByUserId(ctx, userId)
	if err != nil {
		return model.CartItemModel{}, err
	}

	// Find cart item
	var cartItem entity.CartItem
	result := cartService.DB.WithContext(ctx).
		Where("id = ? AND cart_id = ?", cartItemId, cart.Id).
		First(&cartItem)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return model.CartItemModel{}, errors.New("cart item not found")
		}
		return model.CartItemModel{}, result.Error
	}

	// Get product to check stock
	product, err := cartService.ProductRepository.FindByProductId(ctx, cartItem.ProductId)
	if err != nil {
		return model.CartItemModel{}, err
	}

	// Check stock availability
	if product.Stock < request.Quantity {
		return model.CartItemModel{}, errors.New("insufficient stock")
	}

	// Update cart item
	_, err = cartService.CartRepository.UpdateCartItem(ctx, cartItemId, request.Quantity)
	if err != nil {
		return model.CartItemModel{}, err
	}

	// Return updated item
	return model.CartItemModel{
		Id:        cartItemId,
		CartId:    cartItem.CartId,
		ProductId: cartItem.ProductId,
		Quantity:  request.Quantity,
		Price:     cartItem.Price,
		CreatedAt: cartItem.CreatedAt.String(),
		Product: model.ProductModel{
			Id:          strconv.FormatUint(uint64(product.Id), 10),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			ImageUrl:    product.ImageUrl,
		},
	}, nil
}

func (cartService *cartServiceImpl) RemoveFromCart(ctx context.Context, userId uint, cartItemId uint) error {
	// Get cart to verify ownership
	cart, err := cartService.CartRepository.GetCartByUserId(ctx, userId)
	if err != nil {
		return err
	}

	// Find and verify cart item belongs to user's cart
	var cartItem entity.CartItem
	result := cartService.DB.WithContext(ctx).
		Where("id = ? AND cart_id = ?", cartItemId, cart.Id).
		First(&cartItem)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.New("cart item not found")
		}
		return result.Error
	}

	return cartService.CartRepository.RemoveCartItem(ctx, cartItemId)
}

func (cartService *cartServiceImpl) ClearCart(ctx context.Context, userId uint) error {
	cart, err := cartService.CartRepository.GetCartByUserId(ctx, userId)
	if err != nil {
		return err
	}

	return cartService.CartRepository.ClearCart(ctx, cart.Id)
}