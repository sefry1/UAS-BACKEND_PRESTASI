package service

import (
	"prestasi_backend/app/repository"

	"github.com/gofiber/fiber/v2"
)

var lecturerRepo = repository.NewLecturerRepository()

func LecturerList(c *fiber.Ctx) error {
	data, _ := lecturerRepo.FindAll()
	return c.JSON(data)
}

func LecturerAdvisees(c *fiber.Ctx) error {
	id := c.Params("id")
	data, _ := repository.NewStudentRepository().FindByAdvisor(id)
	return c.JSON(data)
}
