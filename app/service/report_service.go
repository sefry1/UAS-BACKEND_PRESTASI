package service

import (
	"github.com/gofiber/fiber/v2"
)

// ==========================================
// REPORT STATISTICS
// ==========================================
func ReportStatistics(c *fiber.Ctx) error {
	all, err := AchievementRefRepo.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	statusCount := map[string]int{}
	typeCount := map[string]int{}

	for _, x := range all {
		statusCount[x.Status]++
		typeCount[x.MongoAchievementID]++
	}

	return c.JSON(fiber.Map{
		"status_distribution": statusCount,
		"type_distribution":   typeCount,
	})
}

// ==========================================
// REPORT FOR SINGLE STUDENT
// ==========================================
func ReportStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	// ✅ FIX: gunakan FindByStudentID (method yang benar)
	achievements, err := AchievementRefRepo.FindByStudentID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"student_id":         id,
		"achievement_total":  len(achievements),
		"achievements":       achievements,
	})
}
