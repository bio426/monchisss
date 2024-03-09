package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/core"
)

const (
	JwtSecret          = "mysecret"
	CookieName         = "monchisss_jwt"
	TokenDurationHours = 12
)

type AuthCtl core.Controller

type CtlLoginResponse struct {
	Token      string
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	ExpiryDate time.Time `json:"expiryDate"`
}

func (ctl *AuthCtl) Login(c echo.Context) error {
	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := Service.Login(c.Request().Context(), SvcLoginParams{
		Username: body.Username,
		Password: body.Password,
	})
	if err != nil {
		switch err {
		case ErrUserUnauthorized:
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		case ErrUserInactive:
			return echo.NewHTTPError(http.StatusGone, err)
		default:
			return err
		}
	}

	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    res.Token,
		Expires:  res.ExpiryDate,
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, res)
}

func (ctl *AuthCtl) Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    "",
		Expires:  time.Now().Add(time.Second * -1),
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

var Controller = &AuthCtl{}
