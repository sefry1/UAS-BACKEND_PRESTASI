package service

import (
	"prestasi_backend/app/repository"

	"github.com/gofiber/fiber/v2"
)

var rolePermRepo = repository.NewRolePermissionRepository()

func RolePermissionList(c *fiber.Ctx) error {
	roleID := c.Query("role_id")

	data, err := rolePermRepo.GetPermissions(roleID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}
