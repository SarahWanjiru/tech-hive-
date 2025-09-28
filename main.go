package main

import (
	"github.com/tech-hive/ecommerce/client/restclient"
	"github.com/tech-hive/ecommerce/configuration"
	"github.com/tech-hive/ecommerce/controller"
	_ "github.com/tech-hive/ecommerce/docs"
	"github.com/tech-hive/ecommerce/exception"
	repository "github.com/tech-hive/ecommerce/repository/impl"
	service "github.com/tech-hive/ecommerce/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Tech Hive E-commerce
// @version 1.0.0
// @description Tech Hive E-commerce application using Go Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9999
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description Authorization For JWT
func main() {
	//setup configuration
	config := configuration.New()
	database := configuration.NewDatabase(config)
	redis := configuration.NewRedis(config)

	//repository
		productRepository := repository.NewProductRepositoryImpl(database)
		transactionRepository := repository.NewTransactionRepositoryImpl(database)
		transactionDetailRepository := repository.NewTransactionDetailRepositoryImpl(database)
		userRepository := repository.NewUserRepositoryImpl(database)
		cartRepository := repository.NewCartRepositoryImpl(database)
		orderRepository := repository.NewOrderRepositoryImpl(database)

	//rest client
	httpBinRestClient := restclient.NewHttpBinRestClient()

	//service
		productService := service.NewProductServiceImpl(&productRepository, redis)
		transactionService := service.NewTransactionServiceImpl(&transactionRepository)
		transactionDetailService := service.NewTransactionDetailServiceImpl(&transactionDetailRepository)
		userService := service.NewUserServiceImpl(&userRepository)
		cartService := service.NewCartServiceImpl(&cartRepository, &productRepository, database)
		orderService := service.NewOrderServiceImpl(&orderRepository, &cartRepository, &productRepository, database)
		mpesaService := service.NewMpesaServiceImpl(config, &orderRepository, database)
		seedService := service.NewSeedServiceImpl(&userRepository, &productRepository, database)
		httpBinService := service.NewHttpBinServiceImpl(&httpBinRestClient)

	//controller
		productController := controller.NewProductController(&productService, config)
		transactionController := controller.NewTransactionController(&transactionService, config)
		transactionDetailController := controller.NewTransactionDetailController(&transactionDetailService, config)
		userController := controller.NewUserController(&userService, config)
		cartController := controller.NewCartController(&cartService, config)
		orderController := controller.NewOrderController(&orderService, config)
		mpesaController := controller.NewMpesaController(&mpesaService, config)
		seedController := controller.NewSeedController(&seedService, config)
		httpBinController := controller.NewHttpBinController(&httpBinService)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://127.0.0.1:3000, http://localhost:9999, http://app:9999",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	//routing
		productController.Route(app)
		transactionController.Route(app)
		transactionDetailController.Route(app)
		userController.Route(app)
		cartController.Route(app)
		orderController.Route(app)
		mpesaController.Route(app)
		seedController.Route(app)
		httpBinController.Route(app)

	//swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//health check
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"service": "mini-ecommerce-api",
		})
	})

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
