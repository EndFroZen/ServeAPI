package function_post

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DataPostApi struct {
	Raw string
}

func PostApi(c *fiber.Ctx) error {
	data := new(DataPostApi)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "nobody",
		})
	}
	fileName := randomString(10) + ".json"
	newFilename := fmt.Sprintf("./jsonadds/%s",fileName)
	err := ioutil.WriteFile(newFilename, []byte(data.Raw), 0644)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "ไม่สามารถบันทึกไฟล์ได้",
		})
	}
	fmt.Println("สร้างไฟล์:", fileName)
	return c.JSON(fiber.Map{
		"message":  "บันทึกไฟล์สำเร็จ",
		"filename": fileName,
	})
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
