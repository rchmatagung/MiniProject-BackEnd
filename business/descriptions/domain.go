package descriptions

import (
	"context"
	"time"
)

type Domain struct {
	Id          uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UseCase interface {
	InsertDescription(ctx context.Context, domain Domain) (Domain, error)
	GetAllDescription(ctx context.Context, search string) ([]Domain, error)
	GetDescriptionById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertDescription(ctx context.Context, domain Domain) (Domain, error)
	GetAllDescription(ctx context.Context, search string) ([]Domain, error)
	GetDescriptionById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}