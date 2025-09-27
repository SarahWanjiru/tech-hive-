package controller

import (
	"github.com/RizkiMufrizal/gofiber-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/model"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/service"
	"github.com/gofiber/fiber/v2"
)

func NewSeedController(seedService *service.SeedService, config configuration.Config) *SeedController {
	return &SeedController{SeedService: *seedService, Config: config}
}

type SeedController struct {
	service.SeedService
	configuration.Config
}

func (controller SeedController) Route(app *fiber.App) {
	app.Post("/v1/api/seed/users", controller.SeedUsers)
	app.Post("/v1/api/seed/products", controller.SeedProducts)
	app.Post("/v1/api/seed/all", controller.SeedAll)
}

// SeedUsers godoc
// @Summary Seed sample users
// @Description Create sample users for testing
// @Tags Seed
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/seed/users [post]
func (controller SeedController) SeedUsers(c *fiber.Ctx) error {
	err := controller.SeedService.SeedUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Error seeding users",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Users seeded successfully",
		Data:    nil,
	})
}

// SeedProducts godoc
// @Summary Seed sample products
// @Description Create sample products for testing
// @Tags Seed
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/seed/products [post]
func (controller SeedController) SeedProducts(c *fiber.Ctx) error {
	err := controller.SeedService.SeedProducts(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Error seeding products",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Products seeded successfully",
		Data:    nil,
	})
}

// SeedAll godoc
// @Summary Seed all sample data
// @Description Create all sample data (users and products) for testing
// @Tags Seed
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/seed/all [post]
func (controller SeedController) SeedAll(c *fiber.Ctx) error {
	err := controller.SeedService.SeedAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    500,
			Message: "Error seeding data",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "All data seeded successfully",
		Data:    nil,
	})
}