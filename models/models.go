package models

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
