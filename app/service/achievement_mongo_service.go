package service

import (
    "prestasi_backend/app/model"
    "time"

    "github.com/gofiber/fiber/v2"
)

func AchievementCreate(c *fiber.Ctx) error {
	var req model.AchievementMongo
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Get user ID from JWT token
	userID := c.Locals("user_id").(string)

	// Verify user is a student (not admin or dosen)
	student, err := StudentRepo.FindByUserID(userID)
	if err != nil {
		return c.Status(403).JSON(fiber.Map{
			"error": "Access denied. Only students can create achievements.",
		})
	}

	// Verify student record exists
	if student == nil || student.ID == "" {
		return c.Status(403).JSON(fiber.Map{
			"error": "Student record not found. Please contact administrator.",
		})
	}

	// Set student ID in MongoDB document
	req.StudentID = student.ID
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	// 1. Insert to MongoDB (flexible achievement details)
	mongoID, err := AchievementMongoRepo.Create(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create achievement: " + err.Error()})
	}

	// 2. Insert to PostgreSQL (workflow reference)
	achievementRefID, err := AchievementRefRepo.Create(student.ID, mongoID)
	if err != nil {
		// Rollback: delete from MongoDB if PostgreSQL insert fails
		AchievementMongoRepo.Delete(mongoID)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create achievement reference: " + err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":              "Achievement created successfully",
		"achievement_id":       achievementRefID,
		"mongo_id":             mongoID,
		"status":               "draft",
		"student_id":           student.ID,
	})
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
