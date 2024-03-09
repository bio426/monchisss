package product

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/auth"
	"github.com/bio426/monchisss/internal/core"
)

type AuthCtl core.Controller

type CtlListRow struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Category *string `json:"category,omitempty"`
}
type CtlListResponse struct {
	Rows []CtlListRow `json:"rows"`
}

func (ctl *AuthCtl) List(c echo.Context) error {
	userStore := c.Get(auth.CtxUserStoreKey).(int32)
	res, err := Service.List(c.Request().Context(), userStore)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *AuthCtl) Create(c echo.Context) error {
	body := struct {
		Type     string             `json:"type" validate:"required,oneof=simple variant compound"`
		Name     string             `json:"name" validate:"required"`
		Price    float32            `json:"price"`
		Category int32              `json:"category" validate:"required"`
		Variants []SvcCreateVariant `json:"variants"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return err
	}

	userStore := c.Get(auth.CtxUserStoreKey).(int32)
	err := Service.Create(c.Request().Context(), SvcCreateParams{
		Type:     body.Type,
		Name:     body.Name,
		Price:    body.Price,
		StoreId:  userStore,
		Category: body.Category,
		Variants: body.Variants,
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

type CtlListCategoryRow struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}
type CtlListCategoryResponse struct {
	Rows []CtlListCategoryRow `json:"rows"`
}

func (ctl *AuthCtl) ListCategory(c echo.Context) error {
	userStore := c.Get(auth.CtxUserStoreKey).(int32)
	res, err := Service.ListCategory(c.Request().Context(), userStore)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *AuthCtl) CreateCategory(c echo.Context) error {
	body := struct {
		Name string `json:"name" validate:"required"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	userStore := c.Get(auth.CtxUserStoreKey).(int32)
	err := Service.CreateCategory(c.Request().Context(), SvcCreateCategoryParams{
		Name:    body.Name,
		StoreId: userStore,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}

var Controller = &AuthCtl{}
