package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/middleware"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

func NewOrderController(orderService *service.OrderService, config configuration.Config) *OrderController {
	return &OrderController{OrderService: *orderService, Config: config}
}

type OrderController struct {
	service.OrderService
	configuration.Config
}

func (controller OrderController) Route(app *fiber.App) {
	// All order routes require authentication
	app.Post("/v1/api/orders", middleware.AuthenticateJWT("customer", controller.Config), controller.CreateOrder)
	app.Get("/v1/api/orders", middleware.AuthenticateJWT("customer", controller.Config), controller.GetUserOrders)
	app.Get("/v1/api/orders/:id", middleware.AuthenticateJWT("customer", controller.Config), controller.GetOrderById)
	app.Put("/v1/api/orders/:id/status", middleware.AuthenticateJWT("admin", controller.Config), controller.UpdateOrderStatus)
	app.Delete("/v1/api/orders/:id", middleware.AuthenticateJWT("customer", controller.Config), controller.CancelOrder)
}

// CreateOrder godoc
// @Summary Create new order from cart
// @Description Create a new order from the user's cart items
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body model.CreateOrderModel true "Create order request"
// @Success 201 {object} model.GeneralResponse
// @Router /v1/api/orders [post]
// @Security JWT
func (controller OrderController) CreateOrder(c *fiber.Ctx) error {
	var request model.CreateOrderModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdFloat := claims["user_id"].(float64)
	userId := uint(userIdFloat)

	order, err := controller.OrderService.CreateOrder(c.Context(), userId, request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Error creating order",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Order created successfully",
		Data:    order,
	})
}

// GetUserOrders godoc
// @Summary Get user's orders
// @Description Get all orders for the authenticated user
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/orders [get]
// @Security JWT
func (controller OrderController) GetUserOrders(c *fiber.Ctx) error {
	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdFloat := claims["user_id"].(float64)
	userId := uint(userIdFloat)

	orders, err := controller.OrderService.GetOrdersByUserId(c.Context(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Error retrieving orders",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    orders,
	})
}

// GetOrderById godoc
// @Summary Get order by ID
// @Description Get a specific order by its ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/orders/{id} [get]
// @Security JWT
func (controller OrderController) GetOrderById(c *fiber.Ctx) error {
	orderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid order ID",
			Data:    err.Error(),
		})
	}

	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdFloat := claims["user_id"].(float64)
	userId := uint(userIdFloat)

	order, err := controller.OrderService.GetOrderById(c.Context(), uint(orderId), userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    404,
			Message: "Order not found",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    order,
	})
}

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Update the status of an order (admin only)
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param request body model.UpdateOrderStatusModel true "Update order status request"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/orders/{id}/status [put]
// @Security JWT
func (controller OrderController) UpdateOrderStatus(c *fiber.Ctx) error {
	var request model.UpdateOrderStatusModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	orderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid order ID",
			Data:    err.Error(),
		})
	}

	order, err := controller.OrderService.UpdateOrderStatus(c.Context(), uint(orderId), request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Error updating order status",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Order status updated successfully",
		Data:    order,
	})
}

// CancelOrder godoc
// @Summary Cancel order
// @Description Cancel a pending order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/orders/{id} [delete]
// @Security JWT
func (controller OrderController) CancelOrder(c *fiber.Ctx) error {
	orderId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Invalid order ID",
			Data:    err.Error(),
		})
	}

	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdFloat := claims["user_id"].(float64)
	userId := uint(userIdFloat)

	err = controller.OrderService.CancelOrder(c.Context(), uint(orderId), userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Error cancelling order",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Order cancelled successfully",
		Data:    nil,
	})
}