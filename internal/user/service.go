package user

import (
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

var Service = &AuthSvc{}
