package service

import (
	"github.com/gofiber/fiber/v2"
)

// =============================
// LIST STUDENTS
// =============================
func StudentList(c *fiber.Ctx) error {
	data, err := StudentRepo.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// =============================
// STUDENT DETAIL
// =============================
func StudentDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := StudentRepo.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Student not found"})
	}
	return c.JSON(data)
}

// =============================
// STUDENT ACHIEVEMENTS
// =============================
func StudentAchievements(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := AchievementRefRepo.FindByStudentID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

// =============================
// SET ADVISOR
// =============================
func StudentSetAdvisor(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct {
		AdvisorID string `json:"advisor_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := StudentRepo.UpdateAdvisor(id, req.AdvisorID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Advisor updated"})
}
