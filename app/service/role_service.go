package service

import (
	"github.com/gofiber/fiber/v2"
)

func RoleList(c *fiber.Ctx) error {
	data, err := RoleRepo.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}
