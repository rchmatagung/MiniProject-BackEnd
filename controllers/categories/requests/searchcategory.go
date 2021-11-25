package requests

import "miniproject/business/categories"

type SearchCategory struct {
	Category string `json:"category"`
}

func ToDomain(search SearchCategory) categories.Domain {
	return categories.Domain{
		Category: search.Category,
	}
}