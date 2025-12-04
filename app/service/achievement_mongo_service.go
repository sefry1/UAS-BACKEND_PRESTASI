package service

import (
    "prestasi_backend/app/model"
    "time"

    "github.com/gofiber/fiber/v2"
)

func AchievementCreate(c *fiber.Ctx) error {
    var req model.AchievementMongo
    c.BodyParser(&req)

    req.CreatedAt = time.Now()
    req.UpdatedAt = time.Now()

    id, err := AchievementMongoRepo.Create(req)
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

    err := AchievementMongoRepo.Update(id, req)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "Updated"})
}

func AchievementDelete(c *fiber.Ctx) error {
    id := c.Params("id")

    err := AchievementMongoRepo.Delete(id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "Deleted"})
}
