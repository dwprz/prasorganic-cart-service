package repository

import (
	"context"
	"encoding/json"

	"github.com/dwprz/prasorganic-cart-service/src/common/errors"
	"github.com/dwprz/prasorganic-cart-service/src/interface/repository"
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	"github.com/dwprz/prasorganic-cart-service/src/model/entity"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type CartImpl struct {
	db *gorm.DB
}

func NewCart(db *gorm.DB) repository.Cart {
	return &CartImpl{
		db: db,
	}
}

func (c *CartImpl) Create(ctx context.Context, data *dto.CreateCartReq) error {
	err := c.db.WithContext(ctx).Table("carts").Create(data).Error

	if errPG, ok := err.(*pgconn.PgError); ok && errPG.Code == "23505" {
		return &errors.Response{HttpCode: 409, Message: "cart already exists"}
	}

	return err
}

func (c *CartImpl) FindManyByUserId(ctx context.Context, userId string, limit, offset int) (*dto.CartWithCountRes, error) {
	queryRes := new(dto.CartQueryRes)

	query := `
	WITH cte_total_cart AS (
		SELECT COUNT(*) FROM carts WHERE user_id = $1
	),
	cte_cart AS (
		SELECT 
			*
		FROM
			carts
		WHERE
			user_id = $1
		ORDER BY
			user_id DESC
		LIMIT $2 OFFSET $3
	)
	SELECT
		(SELECT * FROM cte_total_cart) AS total_cart,
		(SELECT json_agg(row_to_json(cte_cart.*)) FROM cte_cart) AS cart;
	`

	if err := c.db.WithContext(ctx).Raw(query, userId, limit, offset).Scan(queryRes).Error; err != nil {
		return nil, err
	}

	if len(queryRes.Cart) == 0 {
		return nil, &errors.Response{HttpCode: 404, Message: "cart not found"}
	}

	var cart []entity.Cart
	if err := json.Unmarshal(queryRes.Cart, &cart); err != nil {
		return nil, err
	}

	return &dto.CartWithCountRes{
		Cart:      cart,
		TotalCart: queryRes.TotalCart,
	}, nil
}

func (c *CartImpl) CountByUserId(ctx context.Context, userId string) (totalCart int64, err error) {

	err = c.db.WithContext(ctx).Table("carts").Where("user_id = ?", userId).Count(&totalCart).Error
	return totalCart, err
}

func (c *CartImpl) DeleteItem(ctx context.Context, data *dto.DeleteItemCartReq) error {

	res := c.db.WithContext(ctx).Where("user_id = ? AND product_id = ?", data.UserId, data.ProductId).Delete(&entity.Cart{})
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return &errors.Response{HttpCode: 404, Message: "cart not found"}
	}

	return res.Error
}
