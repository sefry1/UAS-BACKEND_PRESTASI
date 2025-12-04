package service

import (
	"prestasi_backend/app/repository"

	"github.com/gofiber/fiber/v2"
)

var lecturerRepo = repository.NewLecturerRepository()
var studentRepoL = repository.NewStudentRepository()

// =====================================================
// LIST ALL LECTURERS
// =====================================================
func LecturerList(c *fiber.Ctx) error {
	data, err := lecturerRepo.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// =====================================================
// GET LECTURER ADVISEES
// =====================================================
func LecturerAdvisees(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := studentRepoL.FindByAdvisor(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "No advisees found"})
	}

	return c.JSON(data)
}
