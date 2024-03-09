package store

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/core"
)

type AuthCtl core.Controller

type CtlListRow struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Admin     string    `json:"admin"`
	CreatedAt time.Time `json:"createdAt"`
}
type CtlListResponse struct {
	Rows []CtlListRow `json:"rows"`
}

func (ctl *AuthCtl) List(c echo.Context) error {
	res, err := Service.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *AuthCtl) Create(c echo.Context) error {
	body := struct {
		Name          string `json:"name" validate:"required"`
		Token         string `json:"token" validate:"required"`
		OwnerUsername string `json:"ownerUsername" validate:"required"`
		OwnerPassword string `json:"ownerPassword" validate:"required"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := Service.Create(c.Request().Context(), SvcCreateParams{
		Name:          body.Name,
		Token:         body.Token,
		OwnerUsername: body.OwnerUsername,
		OwnerPassword: body.OwnerPassword,
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

type CtlListUsersRow struct {
	Id        int32     `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}
type CtlListUsersResponse struct {
	Rows []CtlListUsersRow `json:"rows"`
}

func (ctl *AuthCtl) ListUsers(c echo.Context) error {
	params := struct {
		Id int32 `param:"id" validate:"required"`
	}{}
	if err := c.Bind(&params); err != nil {
		return err
	}
	if err := c.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := Service.ListUsers(c.Request().Context(), params.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *AuthCtl) CreateUser(c echo.Context) error {
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

	res, err := Service.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

var Controller = &AuthCtl{}
