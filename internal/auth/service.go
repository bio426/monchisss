package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/bio426/monchisss/datasource"
	"github.com/bio426/monchisss/internal/core"
)

type AuthSvc core.Service

var ErrUserUnauthorized = errors.New("Unauthorized user")
var ErrUserInactive = errors.New("Inactive user")

type CustomClaims struct {
	jwt.RegisteredClaims
	// User Id
	Id int32 `json:"id"`
	// User role
	Role string `json:"role"`
	// User store id
	Store int32 `json:"store"`
}

type SvcLoginParams struct {
	Username string
	Password string
}

func (svc *AuthSvc) Login(c context.Context, params SvcLoginParams) (*CtlLoginResponse, error) {
	var (
		id       int32
		username string
		password string
		role     string
		active   bool
		store    int32
	)
	row := datasource.Postgres.QueryRowContext(
		c,
		"select id, username, password, role, active, store from users where username = $1",
		params.Username,
	)
	if err := row.Scan(&id, &username, &password, &role, &active, &store); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserUnauthorized
		}
		return nil, err
	}

	// verifica password
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(params.Password)); err != nil {
		return nil, ErrUserUnauthorized
	}

	// verifica que el usuario tenga un store asignado
	if role != "super" && !active {
		return nil, ErrUserInactive
	}

	// asigna el store del usuario
	var storeId int32
	if role != "super" {
		storeId = store
	}

	expiryDate := time.Now().Add(time.Hour * TokenDurationHours)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Id:    id,
		Role:  role,
		Store: storeId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiryDate),
		},
	})
	token, err := claims.SignedString([]byte(JwtSecret))
	if err != nil {
		return nil, err
	}

	res := &CtlLoginResponse{
		Token:      token,
		Username:   params.Username,
		Role:       role,
		ExpiryDate: expiryDate,
	}

	return res, nil
}

var Service = &AuthSvc{}
