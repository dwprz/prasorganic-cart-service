package dto

type Paging struct {
	TotalData int `json:"total_data"`
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
}

type DataWithPaging[T any] struct {
	Data   T       `json:"data"`
	Paging *Paging `json:"paging"`
}
