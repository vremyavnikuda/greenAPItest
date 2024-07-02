package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
	"io/ioutil"
	"net/http"
)

func MakeAPIRequestWithHeaders(c *fiber.Ctx, url string, body interface{}, method string, responseStruct interface{},
	idInstance, apiTokenInstance string) error {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			pterm.Error.Printf("Внутренняя ошибка сервера: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		pterm.Error.Printf("Внутренняя ошибка сервера: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-IdInstance", idInstance)
	req.Header.Set("X-ApiTokenInstance", apiTokenInstance)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		pterm.Error.Printf("Запрос к API не удался: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		pterm.Error.Printf("Не удалось прочитать ответ API.: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	pterm.Info.Printf("Статус ответа API: %d\n", resp.StatusCode)
	pterm.Info.Printf("Тело ответа API: %s\n", respBody)

	if resp.StatusCode != http.StatusOK {
		pterm.Error.Printf("Получен ответ, отличный от 200: %d \n", resp.StatusCode)
		return c.Status(resp.StatusCode).SendString(fmt.Sprintf("Ответ об ошибке от API: %s\n", respBody))
	}

	err = json.Unmarshal(respBody, responseStruct)
	if err != nil {
		pterm.Error.Printf("Не удалось декодировать ответ API.: %v \n", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(responseStruct)
}
