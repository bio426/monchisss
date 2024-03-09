package wa

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/core"
)

const (
	VerifycationToken = "monchisss-token"
)

type WaCtl core.Controller

func (ctl *WaCtl) Verify(c echo.Context) error {
	// this should be validated with path params for each store
	mode := c.QueryParam("hub.mode")
	token := c.QueryParam("hub.verify_token")
	challenge := c.QueryParam("hub.challenge")
	if mode != "subscribe" || token != VerifycationToken {
		return c.NoContent(http.StatusForbidden)
	}

	return c.String(http.StatusOK, challenge)
}

type WhPayload struct {
	Object string    `json:"object"`
	Entry  []WhEntry `json:"entry"`
}
type WhEntry struct {
	Id      string     `json:"id"`
	Changes []WhChange `json:"changes"`
}
type WhChange struct {
	Field string  `json:"field"`
	Value WhValue `json:"value"`
}

// for optional fields we use pointers
type WhValue struct {
	MessagingProduct string     `json:"messaging_product"`
	Metadata         WhMetadata `json:"metadata"`
	// fields below can vary
	Messages []WhMessage `json:"messages"`
}
type WhMetadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberId      string `json:"phone_number_id"`
}
type WhStatus struct {
	Status string `json:"status"`
}
type WhMessage struct {
	From      string `json:"from"`
	Id        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	// fields below can vary
	Text     WhMessageText     `json:"text"`
	Location WhMessageLocation `json:"location"`
}
type WhMessageText struct {
	Body string `json:"body"`
}
type WhMessageLocation struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

var ErrUnknownPayload = errors.New("Unknown payload")

func (ctl *WaCtl) Hook(c echo.Context) error {
	// print body
	if false {
		rBody := echo.Map{}
		err := c.Bind(&rBody)
		if err != nil {
			c.Logger().Error(err)
		}
		jsonData, err := json.Marshal(rBody)
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Println(string(jsonData))
		return c.NoContent(http.StatusOK)
	}
	body := WhPayload{}
	if err := c.Bind(&body); err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusOK)
	}
	if body.Entry == nil || body.Entry[0].Changes == nil || body.Entry[0].Changes[0].Value.Messages == nil {
		c.Logger().Error(ErrUnknownPayload)
		return c.NoContent(http.StatusOK)
	}
	valueObject := body.Entry[0].Changes[0].Value
	message := valueObject.Messages[0]

	payload := SvcProcessMessagePayload{}
	if message.Type == "text" {
		payload.Text = message.Text.Body
	}
	if message.Type == "location" {
		payload.Latitude = message.Location.Latitude
		payload.Longitude = message.Location.Longitude
	}

	err := Service.ProccessMessage(c.Request().Context(), SvcProcessMessageParams{
		From:    message.From,
		To:      valueObject.Metadata.DisplayPhoneNumber,
		Type:    message.Type,
		Payload: payload,
	})
	if err != nil {
		c.Logger().Error(err)
	}

	return c.NoContent(http.StatusOK)
}

var Controller = &WaCtl{}
