package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/common"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/middleware"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

func NewCartController(cartService *service.CartService, config configuration.Config) *CartController {
	return &CartController{CartService: *cartService, Config: config}
}

type CartController struct {
	service.CartService
	configuration.Config
}

func (controller CartController) Route(app *fiber.App) {
	// All cart routes require authentication
	app.Get("/v1/api/cart", middleware.AuthenticateJWT("customer", controller.Config), controller.GetCart)
	app.Post("/v1/api/cart/items", middleware.AuthenticateJWT("customer", controller.Config), controller.AddToCart)
	app.Put("/v1/api/cart/items/:id", middleware.AuthenticateJWT("customer", controller.Config), controller.UpdateCartItem)
	app.Delete("/v1/api/cart/items/:id", middleware.AuthenticateJWT("customer", controller.Config), controller.RemoveFromCart)
	app.Delete("/v1/api/cart", middleware.AuthenticateJWT("customer", controller.Config), controller.ClearCart)
}

// GetCart godoc
// @Summary Get user's cart
// @Description Get the current user's shopping cart
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/cart [get]
// @Security JWT
func (controller CartController) GetCart(c *fiber.Ctx) error {
	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdFloat := claims["user_id"].(float64)
	userId := uint(userIdFloat)

	cart, err := controller.CartService.GetCart(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Error",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    cart,
	})
}

// AddToCart godoc
// @Summary Add item to cart
// @Description Add a product to the user's shopping cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param request body model.AddToCartModel true "Add to cart request"
// @Success 201 {object} model.GeneralResponse
// @Router /v1/api/cart/items [post]
// @Security JWT
func (controller CartController) AddToCart(c *fiber.Ctx) error {
	var request model.AddToCartModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdFloat := claims["user_id"].(float64)
	userId := uint(userIdFloat)

	cartItem, err := controller.CartService.AddToCart(c.Context(), userId, request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Error",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Item added to cart successfully",
		Data:    cartItem,
	})
}

// UpdateCartItem godoc
// @Summary Update cart item quantity
// @Description Update the quantity of an item in the user's cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path int true "Cart Item ID"
// @Param request body model.UpdateCartItemModel true "Update cart item request"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/cart/items/{id} [put]
// @Security JWT
func (controller CartController) UpdateCartItem(c *fiber.Ctx) error {
	var request model.UpdateCartItemModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	cartItemId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid cart item ID",
			Data:    err.Error(),
		})
	}

	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdFloat := claims["user_id"].(float64)
	userId := uint(userIdFloat)

	cartItem, err := controller.CartService.UpdateCartItem(c.Context(), userId, uint(cartItemId), request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Error",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Cart item updated successfully",
		Data:    cartItem,
	})
}

// RemoveFromCart godoc
// @Summary Remove item from cart
// @Description Remove an item from the user's shopping cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path int true "Cart Item ID"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/cart/items/{id} [delete]
// @Security JWT
func (controller CartController) RemoveFromCart(c *fiber.Ctx) error {
	cartItemId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid cart item ID",
			Data:    err.Error(),
		})
	}

	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdFloat := claims["user_id"].(float64)
	userId := uint(userIdFloat)

	err = controller.CartService.RemoveFromCart(c.Context(), userId, uint(cartItemId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Error",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Item removed from cart successfully",
		Data:    nil,
	})
}

// ClearCart godoc
// @Summary Clear cart
// @Description Remove all items from the user's shopping cart
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/cart [delete]
// @Security JWT
func (controller CartController) ClearCart(c *fiber.Ctx) error {
	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdFloat := claims["user_id"].(float64)
	userId := uint(userIdFloat)

	err := controller.CartService.ClearCart(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Error",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Cart cleared successfully",
		Data:    nil,
	})
}