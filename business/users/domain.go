package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	Name      string
	Email     string
	Address   string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	Login(ctx context.Context, email string, password string) (Domain, string, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	GetAllUsers(ctx context.Context) ([]Domain, error)
	GetUserById(ctx context.Context, domain Domain, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	Login(ctx context.Context, domain Domain) (Domain, string, error)
	Register(ctx context.Context, domain Domain) (Domain, error)
	GetAllUsers(ctx context.Context) ([]Domain, error)
	GetUserById(ctx context.Context, id uint) (Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}