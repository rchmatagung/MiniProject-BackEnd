package responses

import (
	"miniproject/business/categories"
	"time"
)

type CategoryResponse struct {
	Id        uint      `json:"id"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updateat"`
}

type SearchResponse struct {
	Category interface{}
}

func FromDomain(domain categories.Domain) CategoryResponse {
	return CategoryResponse{
		Id:        domain.Id,
		Category:  domain.Category,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromCategoriesAll(domain []categories.Domain) []CategoryResponse {
	var all []CategoryResponse
	for _, v := range domain {
		all = append(all, FromDomain(v))
	}
	return all
}