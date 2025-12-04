package service

import (
	"prestasi_backend/app/repository"

	"github.com/gofiber/fiber/v2"
)

var userRepo = repository.NewUserRepository()

// =====================================================
// LIST USERS
// =====================================================
func UserList(c *fiber.Ctx) error {
	data, err := userRepo.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// =====================================================
// USER DETAIL
// =====================================================
func UserDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := userRepo.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(data)
}

// =====================================================
// CREATE USER
// =====================================================
func UserCreate(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
		RoleID   string `json:"role_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := userRepo.Create(req.Username, req.Email, req.Password, req.FullName, req.RoleID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User created"})
}

// =====================================================
// UPDATE USER
// =====================================================
func UserUpdate(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct {
		Email    string `json:"email"`
		FullName string `json:"full_name"`
		RoleID   string `json:"role_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := userRepo.Update(id, req.Email, req.FullName, req.RoleID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User updated"})
}

// =====================================================
// DELETE USER
// =====================================================
func UserDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := userRepo.Delete(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User deleted"})
}

// =====================================================
// UPDATE USER ROLE
// =====================================================
func UserUpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct {
		RoleID string `json:"role_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := userRepo.UpdateRole(id, req.RoleID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Role updated"})
}
