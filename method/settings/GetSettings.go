package settings

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
	"greenAPItest/api"
	"greenAPItest/models"
)

const (
	apiUrl           = "https://1103.api.green-api.com"
	sendFileByUrlAPI = "https://1103.api.green-api.com"
)

func GetSettings(c *fiber.Ctx) error {
	idInstance := c.Get("X-IdInstance")
	apiTokenInstance := c.Get("X-ApiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		pterm.Error.Println("Несанкционированный запрос: отсутствует idInstance или apiTokenInstance.")
		return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует заголовок IdInstance или ApiTokenInstance..")
	}

	fullUrl := fmt.Sprintf("%s/waInstance%s/getSettings/%s", apiUrl, idInstance, apiTokenInstance)
	return api.MakeAPIRequest(c, fullUrl, nil, "GET", &models.SettingsResponse{})
}
