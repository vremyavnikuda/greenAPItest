package settings

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
	"greenAPItest/api"
	"greenAPItest/models"
	"os"
)

func SendMessage(c *fiber.Ctx) error {
	idInstance := c.Get("X-IdInstance")
	apiTokenInstance := c.Get("X-ApiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		pterm.Error.Println("Несанкционированный запрос: отсутствует idInstance или apiTokenInstance.")
		return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует заголовок IdInstance или ApiTokenInstance.")
	}

	var body struct {
		ChatID  string `json:"chatId"`
		Message string `json:"message"`
	}

	if err := c.BodyParser(&body); err != nil {
		pterm.Error.Printf("Неверный запрос: %v\n", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if body.ChatID == "" || body.Message == "" {
		return c.Status(fiber.StatusBadRequest).SendString("ChatId и сообщение являются обязательными полями.")
	}

	apiUrl := os.Getenv("API_URL")
	fullUrl := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", apiUrl, idInstance, apiTokenInstance)
	pterm.Info.Println(fullUrl)
	pterm.Info.Println(body)
	return api.MakeAPIRequest(c, fullUrl, body, "POST", &models.SendMessageResponse{})
}
