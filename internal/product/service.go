package product

import (
	"context"

	"github.com/bio426/monchisss/datasource"
	"github.com/bio426/monchisss/internal/core"
)

type AuthSvc core.Service

func (svc *AuthSvc) List(c context.Context, storeId int32) (CtlListResponse, error) {
	rows, err := datasource.Postgres.QueryContext(
		c,
		`
        select 
          p.id, 
          p.name, 
          p.price, 
          pc.name 
        from 
          products p 
          left join product_categories pc on pc.id = p.category 
        where 
          p.store = $1;
        `,
		storeId,
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
			&row.Price,
			&row.Category,
		); err != nil {
			return CtlListResponse{}, err
		}
		res.Rows = append(res.Rows, row)
	}
	return res, nil
}

type SvcCreateVariant struct {
	Name  string
	Price float32
}
type SvcCreateParams struct {
	Type     string
	Name     string
	Price    float32
	Category int32
	Variants []SvcCreateVariant
	StoreId  int32
}

func (svc *AuthSvc) Create(c context.Context, params SvcCreateParams) error {
	tx, err := datasource.Postgres.BeginTx(c, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var productId int32
	row := tx.QueryRowContext(c,
		"insert into products(type,name,price,store,category) values ($1,$2,$3,$4,$5) returning id",
		params.Type,
		params.Name,
		params.Price,
		params.StoreId,
		params.Category,
	)
	if err := row.Scan(&productId); err != nil {
		return err
	}

	if params.Type == "variant" {
		stmt, err := tx.PrepareContext(c, "insert into product_variants(name,price,product) values ($1,$2,$3)")
		if err != nil {
			return err
		}
		for _, variant := range params.Variants {
			_, err = stmt.ExecContext(c, variant.Name, variant.Price, productId)
			if err != nil {
				return err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (svc *AuthSvc) ListCategory(c context.Context, storeId int32) (CtlListCategoryResponse, error) {
	rows, err := datasource.Postgres.QueryContext(
		c,
		"select id,name from product_categories where store = $1",
		storeId,
	)
	if err != nil {
		return CtlListCategoryResponse{}, err
	}
	defer rows.Close()

	res := CtlListCategoryResponse{Rows: []CtlListCategoryRow{}}
	for rows.Next() {
		var row = CtlListCategoryRow{}
		if err = rows.Scan(
			&row.Id,
			&row.Name,
		); err != nil {
			return CtlListCategoryResponse{}, err
		}
		res.Rows = append(res.Rows, row)
	}
	return res, nil
}

type SvcCreateCategoryParams struct {
	Name    string
	StoreId int32
}

func (svc *AuthSvc) CreateCategory(c context.Context, params SvcCreateCategoryParams) error {
	_, err := datasource.Postgres.ExecContext(c,
		"insert into product_categories(name,store) values ($1,$2)",
		params.Name,
		params.StoreId,
	)
	if err != nil {
		return err
	}

	return nil
}

var Service = &AuthSvc{}
