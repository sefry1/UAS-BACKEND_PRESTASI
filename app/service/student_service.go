package service

import (
	"prestasi_backend/app/repository"
	"github.com/gofiber/fiber/v2"
)

var studentRepo = repository.NewStudentRepository()
var achRefRepo = repository.NewAchievementReferenceRepository()

func StudentList(c *fiber.Ctx) error {
	data, _ := studentRepo.FindAll()
	return c.JSON(data)
}

func StudentDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := studentRepo.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Student not found"})
	}
	return c.JSON(data)
}

func StudentAchievements(c *fiber.Ctx) error {
	id := c.Params("id")
	data, _ := achRefRepo.FindByStudentID(id)
	return c.JSON(data)
}

func StudentSetAdvisor(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct {
		AdvisorID string `json:"advisor_id"`
	}

	c.BodyParser(&req)
	err := studentRepo.UpdateAdvisor(id, req.AdvisorID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Advisor updated"})
}
