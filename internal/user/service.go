package user

import (
	"context"

	"github.com/bio426/monchisss/datasource"
	"github.com/bio426/monchisss/internal/core"
	"golang.org/x/crypto/bcrypt"
)

type UserSvc core.Service

func (svc *UserSvc) List(c context.Context) (CtlListResponse, error) {
	rows, err := datasource.Postgres.QueryContext(
		c,
		"select id,username,role,active,created_at from users where role != 'super'",
	)
	if err != nil {
		return CtlListResponse{}, err
	}
	defer rows.Close()

	res := CtlListResponse{Rows: []CtlListRow{}}
	for rows.Next() {
		var row = CtlListRow{}
		if err = rows.Scan(
			&row.Id,
			&row.Username,
			&row.Role,
			&row.Active,
			&row.CreatedAt,
		); err != nil {
			return CtlListResponse{}, err
		}
		res.Rows = append(res.Rows, row)
	}
	return res, nil
}

type SvcCreateParams struct {
	Username string
	Password string
	Role     string
}

func (svc *UserSvc) Create(c context.Context, params SvcCreateParams) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		return err
	}
	_, err = datasource.Postgres.ExecContext(
		c,
		"insert into users(username,password,role) values ($1,$2,$3)",
		params.Username, hashedPassword, params.Role,
	)
	if err != nil {
		return err
	}
	return nil
}

func (svc *UserSvc) ListInactiveOwners(c context.Context) (CtlInactiveAdminsResponse, error) {
	rows, err := datasource.Postgres.QueryContext(
		c,
		"select id, username from users where active = false and role = 'owner'",
	)
	if err != nil {
		return CtlInactiveAdminsResponse{}, err
	}
	defer rows.Close()

	res := CtlInactiveAdminsResponse{Rows: []CtlInactiveAdminsRow{}}

	for rows.Next() {
		var row = CtlInactiveAdminsRow{}
		if err = rows.Scan(
			&row.Id,
			&row.Username,
		); err != nil {
			return CtlInactiveAdminsResponse{}, err
		}
		res.Rows = append(res.Rows, row)
	}
	return res, nil
}

var Service = &UserSvc{}
