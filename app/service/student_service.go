package service

import (
	"prestasi_backend/app/repository"

	"github.com/gofiber/fiber/v2"
)

var studentRepo = repository.NewStudentRepository()
var achievementRefRepo = repository.NewAchievementReferenceRepository()

// =============================
// LIST STUDENTS
// =============================
func StudentList(c *fiber.Ctx) error {
	data, err := studentRepo.FindAll()
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

	data, err := studentRepo.FindByID(id)
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

	data, err := achievementRefRepo.FindByStudentID(id)
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

	if err := studentRepo.UpdateAdvisor(id, req.AdvisorID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Advisor updated"})
}
