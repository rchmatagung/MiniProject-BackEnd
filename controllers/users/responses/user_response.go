package responses

import (
	"miniproject/business/users"
	"time"
)

type UserResponse struct {
	Id        uint 		`json:"id"`
	Name      string 	`json:"name"`
	Email     string 	`json:"email"`
	Address   string 	`json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LoginResponse struct {
	SessionToken 	string
	User 			interface{}
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:			domain.Id,
		Name:		domain.Name,
		Email:     	domain.Email,
		Address:   	domain.Address,
		CreatedAt: 	domain.CreatedAt,
		UpdatedAt: 	domain.UpdatedAt,
	}
}

func FromUsersAll(domain []users.Domain) []UserResponse {
	var all []UserResponse
	for _, v := range domain {
		all = append(all, FromDomain(v))
	}
	return all
}