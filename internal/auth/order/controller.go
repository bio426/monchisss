package order

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/core"
)

type OrderCtl core.Controller

func (ctl OrderCtl) ConfirmWebhook(c echo.Context) error {
	// i should have this from the meta app
	const myToken = ""

	mode := c.QueryParam("hub.mode")
	token := c.QueryParam("hub.verify_token")
	challenge := c.QueryParam("hub.challenge")

	if mode != "subscribe" || token != myToken {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.String(http.StatusOK, challenge)
}

func (ctl OrderCtl) SendMessage(c echo.Context) error {

	body := struct {
		Product string `json:"messaging_product"`
		RType   string `json:"recipient_type"`
		To      string `json:"to"`
		Type    string `json:"type"`
		Text    string `json:"text"`
	}{
		Product: "whatsapp",
		RType:   "individual",
		To:      "+51910831282",
		Type:    "text",
		Text:    "klk",
	}
	buf := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	http.NewRequest(http.MethodPost, "", buf)
	return c.NoContent(http.StatusOK)
}

var Controller = OrderCtl{}
