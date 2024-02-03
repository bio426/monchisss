package auth

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	// service
	var (
		id       int32
		username string
		password string
		role     string
		active   bool
	)
	row := datasource.Postgres.QueryRowContext(
		c.Request().Context(),
		"select id, username, password, role, active from users where username = $1",
		body.Username,
	)
	if err := row.Scan(&id, &username, &password, &role, &active); err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return err
	}

	if !active {
		return echo.NewHTTPError(http.StatusGone)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	expiryLimit := time.Now().Add(time.Hour * TokenDurationHours)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(id)),
		Subject:   role,
		ExpiresAt: jwt.NewNumericDate(expiryLimit),
	})
	token, err := claims.SignedString([]byte(JwtSecret))
	if err != nil {
		return err
	}
	// ~service

	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    token,
		Expires:  expiryLimit,
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(cookie)

	res := struct {
		Username   string    `json:"username"`
		Role       string    `json:"role"`
		ExpiryDate time.Time `json:"expiryDate"`
	}{
		Username:   username,
		Role:       role,
		ExpiryDate: expiryLimit,
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *AuthCtl) Logout(c echo.Context) error {

	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    "gaaaa",
		Expires:  time.Now().Add(time.Second * -1),
		HttpOnly: true,
		Path:     "/",
	}
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (ctl *AuthCtl) ListUsers(c echo.Context) error {

	// service
	rows, err := datasource.Postgres.QueryContext(
		c.Request().Context(),
		"select id,username,role,created_at from users where role != 'super'",
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	type row = struct {
		Id        int32     `json:"id"`
		Username  string    `json:"username"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"createdAt"`
	}
	rowSlice := []row{}
	for rows.Next() {
		var row = row{}
		if err = rows.Scan(
			&row.Id,
			&row.Username,
			&row.Role,
			&row.CreatedAt,
		); err != nil {
			return err
		}
		rowSlice = append(rowSlice, row)
	}
	// ~
	res := struct {
		Data []row `json:"data"`
	}{
		Data: rowSlice,
	}

	return c.JSON(http.StatusOK, res)
}

var Controller = &AuthCtl{}
