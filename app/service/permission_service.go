package service

import "github.com/gofiber/fiber/v2"

func PermissionList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "permission list"})
}
