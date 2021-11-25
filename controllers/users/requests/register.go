package requests

import "miniproject/business/users"

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func (user *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		Name: 		user.Name, 
		Email: 		user.Email, 
		Password: 	user.Password,
		Address: 	user.Address,  
	}
}