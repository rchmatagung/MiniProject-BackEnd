package requests

import "miniproject/business/descriptions"

type DescriptionSearch struct {
	Description string `json:"description"`
}

func ToDomain(search DescriptionSearch) descriptions.Domain {
	return descriptions.Domain{
		Description: search.Description,
	}
}