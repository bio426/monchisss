package auth

import (
	"context"
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

type AuthSvc core.Service

type SvcRegisterArgs struct {
	Ctx      echo.Context
	Username string
	Password string
	Role     string
}

func (svc *AuthSvc) Register(args SvcRegisterArgs) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(args.Password), 10)
	if err != nil {
		return err
	}
	_, err = datasource.Postgres.ExecContext(
		args.Ctx.Request().Context(),
		"insert into users(username,password,role) values ($1,$2,$3)",
		args.Username, hashedPassword, args.Role,
	)
	if err != nil {
		return err
	}
	return nil
}

type SvcLoginParam struct {
	Username string
	Password string
}

func (svc *AuthSvc) Login(ctx context.Context, args SvcLoginParam) (string, error) {
	var (
		id       int32
		username string
		password string
		role     string
	)
	row := datasource.Postgres.QueryRowContext(
		ctx,
		"select id, username, password, role from users where username = $1",
		args.Username,
	)
	if err := row.Scan(&id, &username, &password, &role); err != nil {
		if err == sql.ErrNoRows {
			return "", echo.NewHTTPError(http.StatusUnauthorized)
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(args.Password)); err != nil {
		return "", echo.NewHTTPError(http.StatusUnauthorized)
	}

	expiryLimit := time.Now().Add(time.Hour * TokenDurationHours)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(id)),
		Subject:   role,
		ExpiresAt: jwt.NewNumericDate(expiryLimit),
	})
	token, err := claims.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

var Service = &AuthSvc{}
