package store

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/bio426/monchisss/datasource"
	"github.com/bio426/monchisss/internal/core"
)

type StoreSvc core.Service

func (svc *StoreSvc) List(c context.Context) (*CtlListResponse, error) {
	rows, err := datasource.Postgres.QueryContext(
		c,
		"select s.id,s.name,s.created_at,u.username from stores s join users u on u.store = s.id where u.role != 'super'",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := &CtlListResponse{Rows: []CtlListRow{}}
	for rows.Next() {
		var row = CtlListRow{}
		if err = rows.Scan(
			&row.Id,
			&row.Name,
			&row.CreatedAt,
			&row.Admin,
		); err != nil {
			return nil, err
		}
		res.Rows = append(res.Rows, row)
	}
	return res, nil
}

type SvcCreateParams struct {
	Name          string
	Token         string
	OwnerUsername string
	OwnerPassword string
}

func (svc *StoreSvc) Create(c context.Context, params SvcCreateParams) error {
	tx, err := datasource.Postgres.BeginTx(c, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// insert store
	var storeId int32
	row := tx.QueryRowContext(c,
		"insert into stores(name,wa_token) values ($1,$2) returning id",
		params.Name,
		params.Token,
	)
	if err := row.Scan(&storeId); err != nil {
		return err
	}

	// insert owner user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.OwnerPassword), 10)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(
		c,
		"insert into users(username,password,role,store) values ($1,$2,$3,$4)",
		params.OwnerUsername, hashedPassword, "owner", storeId,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (svc *StoreSvc) ListUsers(c context.Context, storeId int32) (*CtlListUsersResponse, error) {
	rows, err := datasource.Postgres.QueryContext(
		c,
		"select id,username,role,active,created_at from users where role != 'super' and store = $1",
		storeId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := &CtlListUsersResponse{}
	for rows.Next() {
		var row = CtlListUsersRow{}
		if err = rows.Scan(
			&row.Id,
			&row.Username,
			&row.Role,
			&row.Active,
			&row.CreatedAt,
		); err != nil {
			return nil, err
		}
		res.Rows = append(res.Rows, row)
	}
	return res, nil
}

type SvcCreateUserParams struct {
	Username string
	Password string
	Role     string
}

func (svc *StoreSvc) CreateUser(c context.Context, params SvcCreateUserParams) error {
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

var Service = &StoreSvc{}
