package responses

import (
	"miniproject/business/descriptions"
	"time"
)

type DescriptionResponse struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdat "`
	UpdatedAt   time.Time `json:"updateat "`
}

type SearchResponse struct {
	Description interface{}
}

func FromDomain(domain descriptions.Domain) DescriptionResponse {
	return DescriptionResponse{
		Id:          domain.Id,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromDescriptionsAll(domain []descriptions.Domain) []DescriptionResponse {
	var all []DescriptionResponse
	for _, v := range domain {
		all = append(all, FromDomain(v))
	}
	return all
}