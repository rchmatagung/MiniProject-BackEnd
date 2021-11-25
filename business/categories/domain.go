package categories

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	InsertCategory(ctx context.Context, domain Domain) (Domain, error)
	GetAllCategory(ctx context.Context, search string) ([]Domain, error)
	GetCategoryById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertCategory(ctx context.Context, domain Domain) (Domain, error)
	GetAllCategory(ctx context.Context, search string) ([]Domain, error)
	GetCategoryById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}