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

type SvcCreateParams struct {
	Name  string
	Token string
	Admin int32
}

func (svc *AuthSvc) Create(c context.Context, params SvcCreateParams) error {
	return nil
}

var Service = &AuthSvc{}
