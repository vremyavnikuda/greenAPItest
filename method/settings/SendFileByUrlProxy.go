package settings

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
	"greenAPItest/api"
	"greenAPItest/models"
	"os"
)

func SendFileByUrlProxy(c *fiber.Ctx) error {
	idInstance := c.Get("X-IdInstance")
	apiTokenInstance := c.Get("X-ApiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		pterm.Error.Println("Несанкционированный запрос: отсутствует idInstance или apiTokenInstance.")
		return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует заголовок IdInstance или ApiTokenInstance.")
	}

	var body struct {
		ChatID   string `json:"chatId"`
		URLFile  string `json:"urlFile"`
		FileName string `json:"fileName"`
	}

	if err := c.BodyParser(&body); err != nil {
		pterm.Error.Printf("Неверный запрос: %v\n", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if body.ChatID == "" || body.URLFile == "" || body.FileName == "" {
		return c.Status(fiber.StatusBadRequest).SendString("ChatId, urlFile и fileName являются обязательными полями.")
	}

	sendFileByUrlAPI := os.Getenv("SEND_FILE_BY_URL_API")
	fullUrl := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", sendFileByUrlAPI, idInstance, apiTokenInstance)
	pterm.Info.Println(fullUrl)
	pterm.Info.Println(body)
	return api.MakeAPIRequestWithHeaders(c, fullUrl, body, "POST", &models.SendMessageResponse{}, idInstance, apiTokenInstance)
}
