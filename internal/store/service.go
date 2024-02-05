package store

import (
	"context"

	"github.com/bio426/monchisss/datasource"
	"github.com/bio426/monchisss/internal/core"
)

type AuthSvc core.Service

func (svc *AuthSvc) List(c context.Context) (CtlListResponse, error) {
	rows, err := datasource.Postgres.QueryContext(
		c,
		"select s.id,s.name,s.active,s.created_at,u.username from stores s join users u on u.id = s.admin_id",
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

type SvcCreateParams struct {
	Name  string
	Token string
	Admin int32
}

func (svc *AuthSvc) Create(c context.Context, params SvcCreateParams) error {
	tx, err := datasource.Postgres.BeginTx(c, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(c,
		"insert into stores(name,wa_token,admin_id) values ($1,$2,$3)",
		params.Name,
		params.Token,
		params.Admin,
	)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(
		c,
		"update users set active = $1 where id = $2",
		true,
		params.Admin,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

var Service = &AuthSvc{}
