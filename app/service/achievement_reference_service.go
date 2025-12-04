package service

import (
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

    err := AchievementRefRepo.UpdateStatus(id, "submitted", time.Now(), nil)
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

    file, _ := c.FormFile("file")
    path := "./uploads/" + file.Filename
    c.SaveFile(file, path)

    ach, _ := AchievementMongoRepo.FindByID(id)

    ach.Attachments = append(ach.Attachments, model.Attachment{
        FileName: file.Filename,
        FileURL:  path,
        FileType: file.Header["Content-Type"][0],
    })

    AchievementMongoRepo.Update(id, *ach)

    return c.JSON(fiber.Map{"message": "uploaded"})
}
