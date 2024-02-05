package user

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/core"
)

type UserCtl core.Controller

type CtlListRow struct {
	Id        int32     `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}
type CtlListResponse struct {
	Rows []CtlListRow `json:"rows"`
}

func (ctl *UserCtl) List(c echo.Context) error {
	res, err := Service.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *UserCtl) Create(c echo.Context) error {
	body := struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required,oneof=owner employee"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := Service.Create(c.Request().Context(), SvcCreateParams{
		Username: body.Username,
		Password: body.Password,
		Role:     body.Role,
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

type CtlInactiveAdminsRow struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
}
type CtlInactiveAdminsResponse struct {
	Rows []CtlInactiveAdminsRow `json:"rows"`
}

func (ctl *UserCtl) GetInactiveOwners(c echo.Context) error {
	res, err := Service.ListInactiveOwners(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

var Controller = &UserCtl{}
