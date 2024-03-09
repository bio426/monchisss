package order

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/bio426/monchisss/datasource"
	"github.com/bio426/monchisss/internal/core"
)

type AuthSvc core.Service

var ErrInvalidOrder = errors.New("Invalid order id")

func (svc *AuthSvc) CategoriesById(c context.Context, orderId string) (*CtlCategoriesByIdResponse, error) {
	if orderId != "1234" {
		return nil, ErrInvalidOrder
	}
	storeId := 1
	rows, err := datasource.Postgres.QueryContext(c,
		`
        select 
          p.id, 
          p.type,
          p.name, 
          p.price, 
          pc.name 
        from 
          products p 
          left join product_categories pc on pc.id = p.category 
        where 
          p.store = $1
        `,
		storeId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type rawRow struct {
		id       int32
		Type     string
		name     string
		price    float32
		category string
	}
	rawRows := []rawRow{}
	for rows.Next() {
		var row = rawRow{}
		if err = rows.Scan(
			&row.id,
			&row.Type,
			&row.name,
			&row.price,
			&row.category,
		); err != nil {
			return nil, err
		}
		rawRows = append(rawRows, row)
	}

	// pupulate product variants
	variantIds := []string{}
	for _, row := range rawRows {
		if row.Type == "variant" {
			variantIds = append(variantIds, strconv.Itoa(int(row.id)))
		}
	}
	query := fmt.Sprintf(
		"select pv.id, pv.name, pv.price, p.id from product_variants pv join products p on p.id = pv.product where p.id in (%s)",
		strings.Join(variantIds, ","),
	)
	rows, err = datasource.Postgres.QueryContext(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type rawVariants struct {
		CtlCategoriesByIdItemVariant
		productId int32
	}
	allVariants := []rawVariants{}

	for rows.Next() {
		var row = rawVariants{}
		if err = rows.Scan(
			&row.Id,
			&row.Name,
			&row.Price,
			&row.productId,
		); err != nil {
			return nil, err
		}
		allVariants = append(allVariants, row)
	}

	// construct response
	categories := []CtlCategoriesByIdCategory{}
	for _, row := range rawRows {
		// create item
		item := CtlCategoriesByIdItem{
			Id:    row.id,
			Type:  row.Type,
			Name:  row.name,
			Price: row.price,
		}
		if item.Type == "variant" {
			itemVariants := []CtlCategoriesByIdItemVariant{}
			for _, v := range allVariants {
				if v.productId == item.Id {
					itemVariants = append(itemVariants, CtlCategoriesByIdItemVariant{Id: v.Id, Name: v.Name, Price: v.Price})
				}
			}
			item.Price = 0
			item.Variants = itemVariants
		}

		// if item.Type == "simple" {
		// 	var myPrice float32 = *item.Price
		// 	item.Price = &myPrice
		// }

		// find category to store item
		idx := slices.IndexFunc(categories, func(c CtlCategoriesByIdCategory) bool {
			return c.Name == row.category
		})
		if idx == -1 {
			// if category is not created yet, create it
			category := CtlCategoriesByIdCategory{
				Name:  row.category,
				Items: []CtlCategoriesByIdItem{},
			}
			category.Items = append(category.Items, item)
			categories = append(categories, category)

		} else {
			categories[idx].Items = append(categories[idx].Items, item)
		}
	}

	res := &CtlCategoriesByIdResponse{
		StoreId:    1,
		StoreName:  "Mi tienda",
		StoreImage: "https://placewaifu.com/image/100",
		Categories: categories,
	}

	return res, nil
}

var Service = &AuthSvc{}
