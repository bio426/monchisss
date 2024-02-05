package auth

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/bio426/monchisss/datasource"
	"github.com/bio426/monchisss/internal/core"
)

type AuthSvc core.Service

var ErrUserUnauthorized = errors.New("Unauthorized user")
var ErrUserInactive = errors.New("Inactive user")

type SvcLoginParams struct {
	Username string
	Password string
}

func (svc *AuthSvc) Login(c context.Context, params SvcLoginParams) (CtlLoginResponse, error) {
	var (
		id       int32
		username string
		password string
		role     string
		active   bool
	)
	row := datasource.Postgres.QueryRowContext(
		c,
		"select id, username, password, role, active from users where username = $1",
		params.Username,
	)
	if err := row.Scan(&id, &username, &password, &role, &active); err != nil {
		if err == sql.ErrNoRows {
			return CtlLoginResponse{}, ErrUserUnauthorized
		}
		return CtlLoginResponse{}, err
	}

	if !active {
		return CtlLoginResponse{}, ErrUserInactive
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(params.Password)); err != nil {
		return CtlLoginResponse{}, ErrUserUnauthorized
	}

	expiryDate := time.Now().Add(time.Hour * TokenDurationHours)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(id)),
		Subject:   role,
		ExpiresAt: jwt.NewNumericDate(expiryDate),
	})
	token, err := claims.SignedString([]byte(JwtSecret))
	if err != nil {
		return CtlLoginResponse{}, err
	}

	res := CtlLoginResponse{
		Token:      token,
		Role:       role,
		ExpiryDate: expiryDate,
	}

	return res, nil
}

var Service = &AuthSvc{}
