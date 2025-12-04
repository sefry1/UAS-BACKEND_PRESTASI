package service

import (
	"prestasi_backend/app/repository"

	"github.com/gofiber/fiber/v2"
)

var roleRepo = repository.NewRoleRepository()

func RoleList(c *fiber.Ctx) error {
	data, err := roleRepo.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}
