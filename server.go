package main

import (
	"log"
	function_get "server/function/GET"
	function_post "server/function/POST"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Static("/css", "./css")
	app.Get("/", function_get.Login)
	app.Get("/home",function_get.ServeAPI)
	app.Post("/postapi",function_post.PostApi)
	app.Get("/getapi",function_get.Getapi)
	app.Get("/loadapi",function_get.LoadApi)
	app.Get("/delete",function_post.DeleteApi)
	log.Fatal(app.Listen(":5662"))
}
