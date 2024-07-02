package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/pterm/pterm"
	"greenAPItest/method/settings"
	"os"
)

func main() {
	engine := html.New("./", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("template", nil)
	})

	app.Get("/getSettings", settings.GetSettings)
	app.Get("/getStateInstance", settings.GetStateInstance)
	app.Get("/showMessagesQueue", settings.ShowMessagesQueue)
	app.Get("/clearMessagesQueue", settings.ClearMessagesQueue)
	app.Post("/sendMessage", settings.SendMessage)
	app.Post("/proxy/sendFileByUrl", settings.SendFileByUrlProxy)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	pterm.Info.Printf("Сервер запущен на порту %s\n", port)
	pterm.Fatal.Println(app.Listen(":" + port))
}
