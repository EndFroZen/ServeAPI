package function_get

import (
	"io/ioutil"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func LoadApi(c *fiber.Ctx) error {
	// รับพารามิเตอร์ชื่อไฟล์จาก URL
	fileParam := c.Query("file") // หรือ c.Params("file") ขึ้นกับ route ที่ตั้งไว้

	if fileParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ไม่พบชื่อไฟล์",
		})
	}

	// ป้องกัน path traversal attacks
	cleanFile := filepath.Clean(fileParam)
	// ห้ามใช้ .. เพื่อขึ้นไปโฟลเดอร์อื่น
	if cleanFile != fileParam || cleanFile == "." || cleanFile == ".." {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ชื่อไฟล์ไม่ถูกต้อง",
		})
	}

	filePath := filepath.Join("./jsonadds", cleanFile)

	// อ่านไฟล์
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "ไม่พบไฟล์",
		})
	}

	// ตั้ง content-type เป็น application/json แล้วส่งข้อมูลไฟล์ออกไป
	c.Set("Content-Type", "application/json")
	return c.Send(data)
}
