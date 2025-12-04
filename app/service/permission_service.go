package service

import (
	"prestasi_backend/app/repository"

	"github.com/gofiber/fiber/v2"
)

var permRepo = repository.NewPermissionRepository()

func PermissionList(c *fiber.Ctx) error {
	data, err := permRepo.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}
