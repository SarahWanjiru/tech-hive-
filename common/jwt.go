package common

import (
	"github.com/tech-hive/ecommerce/configuration"
	"github.com/tech-hive/ecommerce/exception"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

func GenerateToken(email string, role string, userId uint, config configuration.Config) string {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	// Create claims with proper structure
	claims := jwt.MapClaims{
		"username": email,  // Use email as username for consistency
		"role":     role,
		"email":    email,
		"user_id":  userId, // Add user ID to claims
		"iat":      time.Now().Unix(),  // Issued at
		"exp":      time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}
