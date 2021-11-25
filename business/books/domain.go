package books

import (
	"context"
	"miniproject/business/categories"
	"miniproject/business/descriptions"
	"time"
)

type Domain struct {
	Id             uint
	Title          string
	Price          uint
	Author         string
	Publisher      string
	Category_Id    uint
	Description_Id uint
	Category       categories.Domain
	Description    descriptions.Domain
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	InsertBook(ctx context.Context, domain *Domain) (Domain, error)
	GetAllBook(ctx context.Context, search string) ([]Domain, error)
	GetBookById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertBook(ctx context.Context, domain *Domain) (Domain, error)
	GetAllBook(ctx context.Context, search string) ([]Domain, error)
	GetBookById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}