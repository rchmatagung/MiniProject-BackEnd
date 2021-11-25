package responses

import "miniproject/business/users"

func UsersLogin(domain users.Domain, token string) LoginResponse {
	response := UserResponse {
		Id:			domain.Id,
		Name:		domain.Name,
		Email:     	domain.Email,
		Address:   	domain.Address,
		CreatedAt: 	domain.CreatedAt,
		UpdatedAt: 	domain.UpdatedAt,
	}

	loginResponse := LoginResponse{}
	loginResponse.SessionToken = token
	loginResponse.User = response
	return loginResponse
}