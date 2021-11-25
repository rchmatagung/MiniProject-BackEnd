package users

import (
	"miniproject/business/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Address   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		Id:			user.Id,
		Name:		user.Name,
		Email:     	user.Email,
		Address:   	user.Address,
		Password:  	user.Password,
		CreatedAt: 	user.CreatedAt,
		UpdatedAt: 	user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:			domain.Id,
		Name:		domain.Name,
		Email:     	domain.Email,
		Address:   	domain.Address,
		Password:  	domain.Password,
		CreatedAt: 	domain.CreatedAt,
		UpdatedAt: 	domain.UpdatedAt,
	}
}

func AllUsers(data []Users) []users.Domain {
	all := []users.Domain{}
	for _, v := range data {
		all = append(all, v.ToDomain())
	}
	return all
}