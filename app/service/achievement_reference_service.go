package service

import (
	"fmt"
	"os"
	"path/filepath"
	"prestasi_backend/app/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

// LIST ALL BY USER
func AchievementList(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	data, err := AchievementRefRepo.FindByUserID(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

// DETAIL
func AchievementDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	ref, err := AchievementRefRepo.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}

	mongoData, _ := AchievementMongoRepo.FindByID(ref.MongoAchievementID)

	return c.JSON(fiber.Map{
		"reference": ref,
		"detail":    mongoData,
	})
}

// SUBMIT
func AchievementSubmit(c *fiber.Ctx) error {
	id := c.Params("id")

	err := AchievementRefRepo.Submit(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Submitted"})
}

// VERIFY
func AchievementVerify(c *fiber.Ctx) error {
	id := c.Params("id")
	verifier := c.Locals("user_id").(string)

	err := AchievementRefRepo.Verify(id, verifier)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Verified"})
}

// REJECT
func AchievementReject(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct{ Reason string `json:"reason"` }
	c.BodyParser(&req)

	verifier := c.Locals("user_id").(string)

	err := AchievementRefRepo.Reject(id, verifier, req.Reason)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Rejected"})
}

// HISTORY
func AchievementHistory(c *fiber.Ctx) error {
	id := c.Params("id")

	ref, _ := AchievementRefRepo.FindByID(id)
	return c.JSON(ref)
}

// UPLOAD FILE
func AchievementUploadAttachment(c *fiber.Ctx) error {
	id := c.Params("id")

	// Get uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No file uploaded"})
	}

	// Create uploads directory if not exists
	uploadsDir := "./uploads"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create upload directory"})
	}

	// Generate unique filename to prevent conflicts
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s_%s", timestamp, file.Filename)
	path := filepath.Join(uploadsDir, filename)
	
	// Save file
	if err := c.SaveFile(file, path); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to save file",
			"detail": err.Error(),
		})
	}

	// Get achievement from MongoDB
	ach, err := AchievementMongoRepo.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Achievement not found"})
	}

	// Add attachment
	ach.Attachments = append(ach.Attachments, model.Attachment{
		FileName: filename,
		FileURL:  path,
		FileType: file.Header.Get("Content-Type"),
	})

	// Update in MongoDB
	if err := AchievementMongoRepo.Update(id, *ach); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update achievement"})
	}

	return c.JSON(fiber.Map{
		"message": "File uploaded successfully",
		"file":    filename,
		"path":    path,
	})
}
