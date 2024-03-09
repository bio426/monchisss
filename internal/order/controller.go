package order

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/core"
)

type AuthCtl core.Controller

type CtlCategoriesByIdItemVariant struct {
	Id    int32   `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
type CtlCategoriesByIdItem struct {
	Id       int32                          `json:"id"`
	Type     string                         `json:"type"`
	Name     string                         `json:"name"`
	Price    float32                        `json:"price,omitempty"`
	Variants []CtlCategoriesByIdItemVariant `json:"variants,omitempty"`
}
type CtlCategoriesByIdCategory struct {
	Name  string                  `json:"name"`
	Items []CtlCategoriesByIdItem `json:"items"`
}
type CtlCategoriesByIdResponse struct {
	StoreId    int32                       `json:"storeId"`
	StoreName  string                      `json:"storeName"`
	StoreImage string                      `json:"storeImage"`
	Categories []CtlCategoriesByIdCategory `json:"categories"`
}

func (ctl *AuthCtl) CategoriesById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	res, err := Service.CategoriesById(c.Request().Context(), id)
	if err != nil {
		if err == ErrInvalidOrder {
			return echo.NewHTTPError(http.StatusGone)
		}
		return err
	}

	return c.JSON(http.StatusOK, res)
}

var Controller = &AuthCtl{}
