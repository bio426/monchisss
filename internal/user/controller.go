package user

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

func (ctl *AuthCtl) List(c echo.Context) error {
	// service
	rows, err := datasource.Postgres.QueryContext(
		c.Request().Context(),
		"select id,username,role,active,created_at from users where role != 'super'",
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	type row = struct {
		Id        int32     `json:"id"`
		Username  string    `json:"username"`
		Role      string    `json:"role"`
		Active    bool      `json:"active"`
		CreatedAt time.Time `json:"createdAt"`
	}
	rowSlice := []row{}
	for rows.Next() {
		var row = row{}
		if err = rows.Scan(
			&row.Id,
			&row.Username,
			&row.Role,
			&row.Active,
			&row.CreatedAt,
		); err != nil {
			return err
		}
		rowSlice = append(rowSlice, row)
	}
	// ~
	res := struct {
		Rows []row `json:"rows"`
	}{
		Rows: rowSlice,
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *AuthCtl) Create(c echo.Context) error {
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

	//
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return err
	}
	_, err = datasource.Postgres.ExecContext(
		c.Request().Context(),
		"insert into users(username,password,role) values ($1,$2,$3)",
		body.Username, hashedPassword, body.Role,
	)
	if err != nil {
		return err
	}
	// ~service

	return c.NoContent(http.StatusOK)
}

var Controller = &AuthCtl{}
