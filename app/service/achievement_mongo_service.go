package service

import (
	"prestasi_backend/app/model"
	"prestasi_backend/app/repository"
	"time"

	"github.com/gofiber/fiber/v2"
)

var achMongoRepo = repository.NewAchievementMongoRepository()

func AchievementCreate(c *fiber.Ctx) error {
	var req model.AchievementMongo
	c.BodyParser(&req)

	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	id, err := achMongoRepo.Create(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"mongo_id": id})
}

func AchievementUpdate(c *fiber.Ctx) error {
	id := c.Params("id")

	var req model.AchievementMongo
	c.BodyParser(&req)

	req.UpdatedAt = time.Now()

	err := achMongoRepo.Update(id, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Updated"})
}

func AchievementDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := achMongoRepo.Delete(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Deleted"})
}
