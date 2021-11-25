package users

import (
	"context"
	"errors"
	"miniproject/business/users"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository (conn *gorm.DB) *UserRepository {
	return &UserRepository {
		Conn: conn,
	}
}

func (repo *UserRepository) Register(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	user := Users {
		Id:			domain.Id,
		Name:		domain.Name,
		Email:     	domain.Email,
		Address:   	domain.Address,
		Password:  	domain.Password,
	}
	result := repo.Conn.Create(&user)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return user.ToDomain(), nil
}

func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]users.Domain, error) {
	var data []Users
	result := repo.Conn.Find(&data)
	if result.Error != nil {
		return []users.Domain{}, result.Error
	}
	return AllUsers(data), nil
}

func (repo *UserRepository) GetUserById(ctx context.Context, id uint) (users.Domain, error) {
	var user Users
	result := repo.Conn.Find(&user, "id = ?", id)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return user.ToDomain(), nil
}

func (repo *UserRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	var user Users
	result := repo.Conn.Find(&user, "email = ?", email)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return user.ToDomain(), nil
}

func (repo *UserRepository) Update(ctx context.Context, domain users.Domain, id uint) (users.Domain, error) {
	data := FromDomain(domain)
	if repo.Conn.Save(&data).Error != nil {
		return users.Domain{}, errors.New("Bad Request")
	}
	return data.ToDomain(), nil
}

func (repo *UserRepository) Delete(ctx context.Context, id uint) error {
	user := Users{}
	err := repo.Conn.Delete(&user, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("Id Not Found")
	}
	return nil
}