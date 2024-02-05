package wa

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/core"
)

const (
	VerifycationToken = "monchisss-token"
)

type WaCtl core.Controller

func (ctl *WaCtl) Verify(c echo.Context) error {
	mode := c.QueryParam("hub.mode")
	token := c.QueryParam("hub.verify_token")
	challenge := c.QueryParam("hub.challenge")
	if mode != "subscribe" || token != VerifycationToken {
		return c.NoContent(http.StatusForbidden)
	}

	return c.String(http.StatusOK, challenge)
}

func (ctl *WaCtl) Hook(c echo.Context) error {
	resBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	c.Logger().Print(string(resBytes))
    // despues de contestar entra en loop, ta raro
	// err = Service.Answer(c.Request().Context())
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	c.Logger().Print("okokok")
	return c.NoContent(http.StatusOK)
}

var Controller = &WaCtl{}
