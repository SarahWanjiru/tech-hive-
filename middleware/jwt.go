package middleware

import (
	"github.com/tech-hive/ecommerce/common"
	"github.com/tech-hive/ecommerce/configuration"
	"github.com/tech-hive/ecommerce/model"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func AuthenticateJWT(role string, config configuration.Config) func(*fiber.Ctx) error {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(jwtSecret),
		TokenLookup:   "header:Authorization",
		AuthScheme:    "Bearer",
		ContextKey:    "user",
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			// Extract user info from token
			userEmail := claims["username"].(string)
			userRole := claims["role"].(string)

			common.NewLogger().Info("JWT validated - User: ", userEmail, " Role: ", userRole, " Required: ", role)

			if userRole == role {
				// Store user info in context for controllers to use
				ctx.Locals("user_email", userEmail)
				ctx.Locals("user_role", userRole)
				return ctx.Next()
			}

			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid Role",
				})
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			common.NewLogger().Error("JWT Authentication Error: ", err.Error())
			common.NewLogger().Error("Request Headers: ", c.GetReqHeaders())
			if err.Error() == "Missing or malformed JWT" {
				return c.
					Status(fiber.StatusBadRequest).
					JSON(model.GeneralResponse{
						Code:    400,
						Message: "Bad Request",
						Data:    "Missing or malformed JWT",
					})
			} else {
				return c.
					Status(fiber.StatusUnauthorized).
					JSON(model.GeneralResponse{
						Code:    401,
						Message: "Unauthorized",
						Data:    "Invalid or expired JWT",
					})
			}
		},
	})
}
