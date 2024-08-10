package helper

import (
	"math"

	"github.com/dwprz/prasorganic-cart-service/src/model/dto"
)

func CreateLimitAndOffset(page int) (limit, offset int) {
	limit = 20
	offset = (page - 1) * limit

	return limit, offset
}

func FormatPagedData[T any](data T, totalData int, page int, limit int) *dto.DataWithPaging[T] {

	return &dto.DataWithPaging[T]{
		Data: data,
		Paging: &dto.Paging{
			TotalData: totalData,
			Page:      page,
			TotalPage: int(math.Ceil(float64(totalData) / float64(limit))),
		},
	}
}
