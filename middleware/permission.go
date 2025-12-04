package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func PermissionRequired(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		perms, ok := c.Locals("permissions").([]interface{})
		if !ok {
			return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
		}

		// cek apakah user punya permission itu
		for _, p := range perms {
			if p.(string) == permission {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{"error": "You do not have permission"})
	}
}
