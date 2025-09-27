package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/middleware"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewMpesaController(mpesaService *service.MpesaService, config configuration.Config) *MpesaController {
	return &MpesaController{MpesaService: *mpesaService, Config: config}
}

type MpesaController struct {
	service.MpesaService
	configuration.Config
}

func (controller MpesaController) Route(app *fiber.App) {
	// Payment routes require authentication
	app.Post("/v1/api/mpesa/stkpush", middleware.AuthenticateJWT("customer", controller.Config), controller.InitiateSTKPush)
	app.Post("/v1/api/mpesa/callback", controller.ProcessCallback)
}

// InitiateSTKPush godoc
// @Summary Initiate M-Pesa STK Push
// @Description Initiate M-Pesa payment for an order
// @Tags M-Pesa
// @Accept json
// @Produce json
// @Param request body model.MpesaPaymentRequest true "M-Pesa payment request"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/mpesa/stkpush [post]
// @Security JWT
func (controller MpesaController) InitiateSTKPush(c *fiber.Ctx) error {
	var request model.MpesaPaymentRequest
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdFloat := claims["user_id"].(float64)
	userId := uint(userIdFloat)

	// Verify order belongs to user (we'll need to add a method to check order ownership)
	// For now, we'll assume the order exists and proceed

	if order.UserId != userId {
		return c.Status(fiber.StatusForbidden).JSON(model.GeneralResponse{
			Code:    403,
			Message: "Forbidden",
			Data:    "You don't have permission to pay for this order",
		})
	}

	response, err := controller.MpesaService.InitiateSTKPush(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Payment initiation failed",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "STK push initiated successfully",
		Data:    response,
	})
}

// ProcessCallback godoc
// @Summary Process M-Pesa callback
// @Description Process M-Pesa payment callback (webhook)
// @Tags M-Pesa
// @Accept json
// @Produce json
// @Param request body model.MpesaCallbackRequest true "M-Pesa callback request"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/mpesa/callback [post]
func (controller MpesaController) ProcessCallback(c *fiber.Ctx) error {
	var request model.MpesaCallbackRequest
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	err = controller.MpesaService.ProcessCallback(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Callback processing failed",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Callback processed successfully",
		Data:    nil,
	})
}