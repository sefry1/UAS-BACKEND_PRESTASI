package service

import (
	"prestasi_backend/app/repository"
	"time"

	"github.com/gofiber/fiber/v2"
)

var refRepo = repository.NewAchievementReferenceRepository()

func AchievementList(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	data, _ := refRepo.FindByUserID(userID)
	return c.JSON(data)
}

func AchievementDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	ref, err := refRepo.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}

	mongoData, _ := achMongoRepo.FindByID(ref.MongoAchievementID)

	return c.JSON(fiber.Map{
		"reference": ref,
		"detail":    mongoData,
	})
}

func AchievementSubmit(c *fiber.Ctx) error {
	id := c.Params("id")

	err := refRepo.UpdateStatus(id, "submitted", time.Now(), nil)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Submitted"})
}

func AchievementVerify(c *fiber.Ctx) error {
	id := c.Params("id")
	verifier := c.Locals("user_id").(string)

	err := refRepo.Verify(id, verifier)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Verified"})
}

func AchievementReject(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct{ Reason string `json:"reason"` }
	c.BodyParser(&req)

	verifier := c.Locals("user_id").(string)

	err := refRepo.Reject(id, verifier, req.Reason)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Rejected"})
}
