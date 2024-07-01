package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/pterm/pterm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	apiUrl   = "https://1103.api.green-api.com"
	apiMedia = "https://1103.media.green-api.com"
)

type SettingsResponse struct {
	Wid                               string `json:"wid"`
	CountryInstance                   string `json:"countryInstance"`
	TypeAccount                       string `json:"typeAccount"`
	WebhookUrl                        string `json:"webhookUrl"`
	WebhookUrlToken                   string `json:"webhookUrlToken"`
	DelaySendMessagesMilliseconds     int    `json:"delaySendMessagesMilliseconds"`
	MarkIncomingMessagesReaded        string `json:"markIncomingMessagesReaded"`
	MarkIncomingMessagesReadedOnReply string `json:"markIncomingMessagesReadedOnReply"`
	SharedSession                     string `json:"sharedSession"`
	ProxyInstance                     string `json:"proxyInstance"`
	OutgoingWebhook                   string `json:"outgoingWebhook"`
	OutgoingMessageWebhook            string `json:"outgoingMessageWebhook"`
	OutgoingAPIMessageWebhook         string `json:"outgoingAPIMessageWebhook"`
	IncomingWebhook                   string `json:"incomingWebhook"`
	DeviceWebhook                     string `json:"deviceWebhook"`
	StatusInstanceWebhook             string `json:"statusInstanceWebhook"`
	StateWebhook                      string `json:"stateWebhook"`
	EnableMessagesHistory             string `json:"enableMessagesHistory"`
	KeepOnlineStatus                  string `json:"keepOnlineStatus"`
	PollMessageWebhook                string `json:"pollMessageWebhook"`
	IncomingBlockWebhook              string `json:"incomingBlockWebhook"`
	IncomingCallWebhook               string `json:"incomingCallWebhook"`
}

type StateInstanceResponse struct {
	StateInstance string `json:"stateInstance"`
}

type SendMessageResponse struct {
	IDMessage string `json:"idMessage"`
}

type ShowMessagesQueueResponse struct {
	MessageID   string   `json:"messageID,omitempty"`
	MessagesIDs []string `json:"messagesIDs,omitempty"`
	Type        string   `json:"type"`
	Body        struct {
		ChatID          string   `json:"chatId"`
		Message         string   `json:"message,omitempty"`
		Messages        []string `json:"messages,omitempty"`
		LinkPreview     bool     `json:"linkPreview,omitempty"`
		QuotedMessageID string   `json:"quotedMessageId,omitempty"`
		Options         []struct {
			OptionName string `json:"optionName"`
		} `json:"options,omitempty"`
		FileName     string `json:"fileName,omitempty"`
		Caption      string `json:"caption,omitempty"`
		URLFile      string `json:"urlFile,omitempty"`
		Latitude     string `json:"latitude,omitempty"`
		Longitude    string `json:"longitude,omitempty"`
		NameLocation string `json:"nameLocation,omitempty"`
		Address      string `json:"address,omitempty"`
		Contact      struct {
			PhoneContact string `json:"phoneContact"`
			FirstName    string `json:"firstName"`
			LastName     string `json:"lastName"`
			MiddleName   string `json:"middleName"`
			Company      string `json:"company"`
		} `json:"contact,omitempty"`
		BackgroundColor string   `json:"backgroundColor,omitempty"`
		Font            string   `json:"font,omitempty"`
		Participants    []string `json:"participants,omitempty"`
		URLLink         string   `json:"urlLink,omitempty"`
		ChatIDFrom      string   `json:"chatIdFrom,omitempty"`
	} `json:"body"`
}

type ClearMessagesQueueResponse struct {
	IsCleared bool `json:"isCleared"`
}

func main() {
	engine := html.New("./", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("template", nil)
	})

	app.Get("/getSettings", getSettings)
	app.Get("/getStateInstance", getStateInstance)
	app.Get("/showMessagesQueue", showMessagesQueue)
	app.Get("/clearMessagesQueue", clearMessagesQueue)
	app.Post("/sendMessage", sendMessage)
	app.Post("/sendFileByUrl", sendFileByUrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}

func getSettings(c *fiber.Ctx) error {
	idInstance := c.Get("X-IdInstance")
	apiTokenInstance := c.Get("X-ApiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		log.Println("Несанкционированный запрос: отсутствует idInstance или apiTokenInstance.")
		return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует заголовок IdInstance или ApiTokenInstance.")
	}

	fullUrl := fmt.Sprintf("%s/waInstance%s/getSettings/%s", apiUrl, idInstance, apiTokenInstance)
	return makeAPIRequest(c, fullUrl, nil, "GET", &SettingsResponse{})
}

func getStateInstance(c *fiber.Ctx) error {
	idInstance := c.Get("X-IdInstance")
	apiTokenInstance := c.Get("X-ApiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		log.Println("Несанкционированный запрос: отсутствует idInstance или apiTokenInstance.")
		return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует заголовок IdInstance или ApiTokenInstance.")
	}

	fullUrl := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", apiUrl, idInstance, apiTokenInstance)
	return makeAPIRequest(c, fullUrl, nil, "GET", &StateInstanceResponse{})
}

func showMessagesQueue(c *fiber.Ctx) error {
	idInstance := c.Get("X-IdInstance")
	apiTokenInstance := c.Get("X-ApiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		log.Println("Несанкционированный запрос: отсутствует idInstance или apiTokenInstance.")
		return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует заголовок IdInstance или ApiTokenInstance.")
	}

	fullUrl := fmt.Sprintf("%s/waInstance%s/showMessagesQueue/%s", apiUrl, idInstance, apiTokenInstance)
	return makeAPIRequest(c, fullUrl, nil, "GET", &[]ShowMessagesQueueResponse{})
}

func clearMessagesQueue(c *fiber.Ctx) error {
	idInstance := c.Get("X-IdInstance")
	apiTokenInstance := c.Get("X-ApiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		log.Println("Несанкционированный запрос: отсутствует idInstance или apiTokenInstance.")
		return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует заголовок IdInstance или ApiTokenInstance.")
	}

	fullUrl := fmt.Sprintf("%s/waInstance%s/clearMessagesQueue/%s", apiUrl, idInstance, apiTokenInstance)
	return makeAPIRequest(c, fullUrl, nil, "GET", &ClearMessagesQueueResponse{})
}

func sendMessage(c *fiber.Ctx) error {
	idInstance := c.Get("X-IdInstance")
	apiTokenInstance := c.Get("X-ApiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		log.Println("Несанкционированный запрос: отсутствует idInstance или apiTokenInstance.")
		return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует заголовок IdInstance или ApiTokenInstance.")
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		log.Printf("Bad request: %v\n", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	fullUrl := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", apiUrl, idInstance, apiTokenInstance)
	pterm.Info.Println(fullUrl)
	pterm.Info.Println(body)
	return makeAPIRequest(c, fullUrl, body, "POST", &SendMessageResponse{})
}

func sendFileByUrl(c *fiber.Ctx) error {
	idInstance := c.Get("X-IdInstance")
	apiTokenInstance := c.Get("X-ApiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		log.Println("Несанкционированный запрос: отсутствует idInstance или apiTokenInstance.")
		return c.Status(fiber.StatusUnauthorized).SendString("Отсутствует заголовок IdInstance или ApiTokenInstance.")
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		log.Printf("Bad request: %v\n", err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	fullUrl := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", apiMedia, idInstance, apiTokenInstance)
	return makeAPIRequest(c, fullUrl, body, "POST", &SendMessageResponse{})
}

func makeAPIRequest(c *fiber.Ctx, url string, body interface{}, method string, responseStruct interface{}) error {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			log.Printf("Внутренняя ошибка сервера: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Внутренняя ошибка сервера: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Запрос к API не удался: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Не удалось прочитать ответ API.: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	log.Printf("Статус ответа API: %d", resp.StatusCode)
	log.Printf("Тело ответа API: %s", respBody)

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-200 response: %d\n", resp.StatusCode)
		return c.Status(resp.StatusCode).SendString(fmt.Sprintf("Ответ об ошибке от API: %s", respBody))
	}

	err = json.Unmarshal(respBody, responseStruct)
	if err != nil {
		log.Printf("Не удалось декодировать ответ API.: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(responseStruct)
}
