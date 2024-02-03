package store

import (
	"context"

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

func (svc *AuthSvc) List(c context.Context) (CtlListResponse, error) {
	rows, err := datasource.Postgres.QueryContext(
		c,
		"select s.id,s.name,s.active,s.created_at,u.username from stores s join users u on u.id = s.admin_id",
	)
	if err != nil {
		return CtlListResponse{}, err
	}
	defer rows.Close()

	res := CtlListResponse{}

	for rows.Next() {
		var row = CtlListRow{}
		if err = rows.Scan(
			&row.Id,
			&row.Name,
			&row.Active,
			&row.CreatedAt,
			&row.Admin,
		); err != nil {
			return CtlListResponse{}, err
		}
		res.Rows = append(res.Rows, row)
	}
	return res, nil
}

func (svc *AuthSvc) Create(args SvcRegisterArgs) error {
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
