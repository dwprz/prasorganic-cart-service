package benchmark

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/dwprz/prasorganic-cart-service/src/common/errors"
	"github.com/dwprz/prasorganic-cart-service/src/common/helper"
	"github.com/dwprz/prasorganic-cart-service/src/infrastructure/database"
	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
	"github.com/dwprz/prasorganic-cart-service/src/model/entity"
	"gorm.io/gorm"
)

// *cd current directory
// go test -v -bench=. -count=1 -p=1

var db *gorm.DB

func init() {
	db = database.NewPostgres()
}

func fullCTE(ctx context.Context, userId string, limit, offset int) (*dto.CartWithCountRes, error) {
	var queryRes []*entity.CartQueryRes

	query := `
	WITH cte_total_cart AS (
		SELECT COUNT(*) AS total_cart FROM carts WHERE user_id = $1
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
	SELECT ctc.total_cart, cc.* FROM cte_total_cart AS ctc CROSS JOIN cte_cart AS cc;
	`

	if err := db.WithContext(ctx).Raw(query, userId, limit, offset).Scan(&queryRes).Error; err != nil {
		return nil, err
	}

	if len(queryRes) == 0 {
		return nil, &errors.Response{HttpCode: 404, Message: "cart not found"}
	}

	cart, total := helper.MapCartQueryToEntities(queryRes)

	return &dto.CartWithCountRes{
		Cart:      cart,
		TotalCart: total,
	}, nil
}

type CartQueryRes struct {
	Cart      []byte
	TotalCart int
}

func cteWithJsonAgg(ctx context.Context, userId string, limit, offset int) (*dto.CartWithCountRes, error) {
	queryRes := new(CartQueryRes)

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

	if err := db.WithContext(ctx).Raw(query, userId, limit, offset).Scan(queryRes).Error; err != nil {
		return nil, err
	}

	if len(queryRes.Cart) == 0 {
		return nil, &errors.Response{HttpCode: 404, Message: "cart not found"}
	}

	var cart []*entity.Cart
	if err := json.Unmarshal(queryRes.Cart, &cart); err != nil {
		return nil, err
	}

	return &dto.CartWithCountRes{
		Cart:      cart,
		TotalCart: queryRes.TotalCart,
	}, nil
}

func nonCTE(ctx context.Context, userId string, limit, offset int) (*dto.CartWithCountRes, error) {
	var totalCart int

	if err := db.WithContext(ctx).Raw("SELECT COUNT(*) AS total_cart FROM carts WHERE user_id = $1", userId).Scan(&totalCart).Error; err != nil {
		return nil, err
	}

	var cart []*entity.Cart

	query := `
	SELECT 
		*
	FROM
		carts
	WHERE
		user_id = $1
	ORDER BY
		user_id DESC
	LIMIT $2 OFFSET $3
	`

	if err := db.WithContext(ctx).Raw(query, userId, limit, offset).Scan(&cart).Error; err != nil {
		return nil, err
	}

	if len(cart) == 0 {
		return nil, &errors.Response{HttpCode: 404, Message: "cart not found"}
	}

	return &dto.CartWithCountRes{
		Cart:      cart,
		TotalCart: totalCart,
	}, nil
}

func Benchmark_CompareQueryCTE(b *testing.B) {
	b.Run("Full CTE", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fullCTE(context.Background(), "user_01", 20, 0)
		}
	})

	b.Run("CTE With JSON Agg", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cteWithJsonAgg(context.Background(), "user_01", 20, 0)
		}
	})

	b.Run("Non CTE", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nonCTE(context.Background(), "user_01", 20, 0)
		}
	})
}

// 1 ms = 1.000.000 ns
// 1 s = 1000 ms

//================================ Full CTE ================================

// test 1
// Benchmark_CompareQueryCTE/Full_CTE-12               5755            202402 ns/op
// PASS
// ok      github.com/dwprz/prasorganic-cart-service/src/repository/benchmark      1.200s

// test 2
// Benchmark_CompareQueryCTE/Full_CTE-12               5620            203740 ns/op
// PASS
// ok      github.com/dwprz/prasorganic-cart-service/src/repository/benchmark      1.181s

// test 3
// Benchmark_CompareQueryCTE/Full_CTE-12               6192            203129 ns/op
// PASS
// ok      github.com/dwprz/prasorganic-cart-service/src/repository/benchmark      1.293s

//================================ CTE With JSON Agg ================================

// test 1
// Benchmark_CompareQueryCTE/CTE_With_JSON_Agg-12              5004            218170 ns/op
// PASS
// ok      github.com/dwprz/prasorganic-cart-service/src/repository/benchmark      1.131s

// test 2
// Benchmark_CompareQueryCTE/CTE_With_JSON_Agg-12              4593            223312 ns/op
// PASS
// ok      github.com/dwprz/prasorganic-cart-service/src/repository/benchmark      1.066s

// test 3
// Benchmark_CompareQueryCTE/CTE_With_JSON_Agg-12              5270            220892 ns/op
// PASS
// ok      github.com/dwprz/prasorganic-cart-service/src/repository/benchmark      1.202s

//================================ Non CTE ================================

// test 1
// Benchmark_CompareQueryCTE/Non_CTE-12                4148            289332 ns/op
// PASS
// ok      github.com/dwprz/prasorganic-cart-service/src/repository/benchmark      2.232s

// test 2
// Benchmark_CompareQueryCTE/Non_CTE-12                4056            290806 ns/op
// PASS
// ok      github.com/dwprz/prasorganic-cart-service/src/repository/benchmark      1.224s

// test 3
// Benchmark_CompareQueryCTE/Non_CTE-12                3676            290244 ns/op
// PASS
// ok      github.com/dwprz/prasorganic-cart-service/src/repository/benchmark      1.116s
