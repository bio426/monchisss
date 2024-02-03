package store

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/bio426/monchisss/datasource"
	"github.com/bio426/monchisss/internal/core"
)

type AuthCtl core.Controller

const (
	JwtSecret          = "mysecret"
	CookieName         = "monchisss_jwt"
	TokenDurationHours = 1
)

type CtlListRow struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Admin     string    `json:"admin"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}
type CtlListResponse struct {
	Rows []CtlListRow `json:"rows"`
}

func (ctl *AuthCtl) List(c echo.Context) error {
	res, err := Service.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *AuthCtl) Create(c echo.Context) error {
	body := struct {
		Name  string `json:"name" validate:"required"`
		Token string `json:"token" validate:"required"`
		Admin int32  `json:"admin" validate:"required"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := Service.Create(c.Request().Context(), SvcCreateParams{
		Name:  body.Name,
		Token: body.Token,
		Admin: body.Admin,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

var Controller = &AuthCtl{}
