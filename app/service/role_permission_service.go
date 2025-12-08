package service

import (
	"github.com/gofiber/fiber/v2"
)

func RolePermissionList(c *fiber.Ctx) error {
	roleID := c.Query("role_id")

	data, err := RolePermissionRepo.GetPermissions(roleID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}
