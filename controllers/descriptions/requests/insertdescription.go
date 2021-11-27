package requests

import "miniproject/business/descriptions"

type InsertDescription struct {
	Description string `json:"description"`
}

func (description *InsertDescription) ToDomain() *descriptions.Domain {
	return &descriptions.Domain{
		Description: description.Description,
	}
}