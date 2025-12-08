package service

import (
	"github.com/gofiber/fiber/v2"
)

func PermissionList(c *fiber.Ctx) error {
	data, err := PermissionRepo.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}
