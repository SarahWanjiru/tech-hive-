package controller

import (
	"github.com/tech-hive/ecommerce/common"
	"github.com/tech-hive/ecommerce/configuration"
	"github.com/tech-hive/ecommerce/exception"
	"github.com/tech-hive/ecommerce/model"
	"github.com/tech-hive/ecommerce/service"
	"github.com/gofiber/fiber/v2"
)

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

type UserController struct {
	service.UserService
	configuration.Config
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/api/authentication", controller.Authentication)
	app.Post("/v1/api/users", controller.Register)
}

// Authentication func Authenticate user.
// @Description authenticate user.
// @Summary authenticate user
// @Tags Authenticate user
// @Accept json
// @Produce json
// @Param request body model.UserModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/authentication [post]
func (controller UserController) Authentication(c *fiber.Ctx) error {
	var request model.UserModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := controller.UserService.Authentication(c.Context(), request)
	userRoles := []map[string]interface{}{
		{
			"role": result.Role,
		},
	}
	tokenJwtResult := common.GenerateToken(result.Email, userRoles, controller.Config)
	resultWithToken := map[string]interface{}{
		"token":  tokenJwtResult,
		"email":  result.Email,
		"name":   result.Name,
		"role":   userRoles,
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    resultWithToken,
	})
}

// Register func Register new user.
// @Description register new user.
// @Summary register new user
// @Tags Register user
// @Accept json
// @Produce json
// @Param request body model.UserRegistrationModel true "Request Body"
// @Success 201 {object} model.GeneralResponse
// @Router /v1/api/users [post]
func (controller UserController) Register(c *fiber.Ctx) error {
	var request model.UserRegistrationModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := controller.UserService.Register(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "User registered successfully",
		Data:    result,
	})
}
