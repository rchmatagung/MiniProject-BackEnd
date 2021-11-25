package responses

import (
	"miniproject/business/books"
	"time"
)

type BookResponse struct {
	Id             uint      `json:"id"`
	Title          string    `json:"title"`
	Price          uint      `json:"price"`
	Author         string    `json:"author"`
	Publisher      string    `json:"publisher"`
	Category_Id    uint      `json:"category_id"`
	Description_Id uint      `json:"description_id"`
	CreatedAt      time.Time `json:"createdat "`
	UpdatedAt      time.Time `json:"updateat "`
}
type SearchResponse struct {
	Book interface{}
}

func FromDomain(domain books.Domain) BookResponse {
	return BookResponse{
		Id:             domain.Id,
		Title:          domain.Title,
		Price:          domain.Price,
		Author:         domain.Author,
		Publisher:      domain.Publisher,
		Category_Id:    domain.Category_Id,
		Description_Id: domain.Category_Id,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func FromBookAll(domain []books.Domain) []BookResponse {
	var all []BookResponse
	for _, v := range domain {
		all = append(all, FromDomain(v))
	}
	return all
}