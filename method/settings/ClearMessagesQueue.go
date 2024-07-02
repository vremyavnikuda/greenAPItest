package settings

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
	"greenAPItest/api"
	"greenAPItest/models"
	"os"
)

func ClearMessagesQueue(c *fiber.Ctx) error {
	idInstance := c.Get("X-IdInstance")
	apiTokenInstance := c.Get("X-ApiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		pterm.Error.Println("Несанкционированный запрос: отсутствует idInstance или apiTokenInstance.")
		return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует заголовок IdInstance или ApiTokenInstance.")
	}

	apiUrl := os.Getenv("API_URL")
	fullUrl := fmt.Sprintf("%s/waInstance%s/clearMessagesQueue/%s", apiUrl, idInstance, apiTokenInstance)
	return api.MakeAPIRequest(c, fullUrl, nil, "GET", &models.ClearMessagesQueueResponse{})
}
