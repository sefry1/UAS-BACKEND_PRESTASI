package service

import "github.com/gofiber/fiber/v2"

func RoleList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "role list"})
}
