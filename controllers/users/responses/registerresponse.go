package responses

import (
	"miniproject/business/users"
	"time"
)

type UserRegister struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromUsersRegister(domain users.Domain) UserResponse {
	return UserResponse{
		Id:			domain.Id,
		Name:		domain.Name,
		Email:     	domain.Email,
		Address:   	domain.Address,
		CreatedAt: 	domain.CreatedAt,
		UpdatedAt: 	domain.UpdatedAt,
	}
}