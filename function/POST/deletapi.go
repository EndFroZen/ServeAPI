package function_post

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func DeleteApi(c *fiber.Ctx) error {
	fileParam := c.Query("file")

	if fileParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing 'file' query parameter",
		})
	}

	// โฟลเดอร์ปลอดภัยที่อนุญาตให้ลบไฟล์
	baseDir := "./jsonadds"

	// ทำให้ path สะอาด (เช่น ตัด ../../)
	cleanFile := filepath.Clean(fileParam)
	fullPath := filepath.Join(baseDir, cleanFile)

	// ตรวจสอบว่า path อยู่ใน baseDir จริง ๆ
	if !strings.HasPrefix(filepath.Clean(fullPath), filepath.Clean(baseDir)+string(os.PathSeparator)) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied: invalid file path",
		})
	}

	// ลบไฟล์
	if err := os.Remove(fullPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to delete file",
			"detail": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "File deleted successfully",
		"file":    cleanFile,
	})
}
