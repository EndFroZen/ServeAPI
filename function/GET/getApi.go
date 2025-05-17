package function_get

import (
	"io/ioutil"
	"net/url"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func Getapi(c *fiber.Ctx) error {
	// โฟลเดอร์ที่เก็บไฟล์ json
	dirPath := "./jsonadds"

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "ไม่สามารถอ่านโฟลเดอร์ได้",
		})
	}

	baseURL := "http://127.0.0.1:5662/loadapi?file="
	deleteURL := "http://127.0.0.1:5662/delete?file="

	type FileInfo struct {
		Filename string `json:"filename"`
		Link     string `json:"link"`
		Delete   string`json:"delete"`
	}

	var fileList []FileInfo

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			// URL encode ชื่อไฟล์ด้วยกันปัญหาชื่อไฟล์มีช่องว่าง หรืออักขระพิเศษ
			encodedFile := url.QueryEscape(file.Name())

			fileList = append(fileList, FileInfo{
				Filename: file.Name(),
				Link:     baseURL + encodedFile,
				Delete: deleteURL+encodedFile,
			})
		}
	}

	return c.JSON(fiber.Map{
		"get": fileList,
	})
}
