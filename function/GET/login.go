package function_get

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.Render("home",nil)
}
