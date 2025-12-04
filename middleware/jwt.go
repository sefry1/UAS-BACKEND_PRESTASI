package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(401).JSON(fiber.Map{"error": "Missing or invalid token"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			return c.Status(500).JSON(fiber.Map{"error": "JWT secret not configured"})
		}

		// PARSING TOKEN JWT v5
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		// CLAIMS
		claims := token.Claims.(jwt.MapClaims)

		c.Locals("user_id", claims["id"])
		c.Locals("role_id", claims["role_id"])
		c.Locals("permissions", claims["permissions"])

		return c.Next()
	}
}
