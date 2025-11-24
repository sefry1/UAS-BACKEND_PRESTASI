package service

import "github.com/gofiber/fiber/v2"

func RolePermissionList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "role-permission list"})
}
