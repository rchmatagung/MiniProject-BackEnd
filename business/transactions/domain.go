package transactions

import (
	"context"
	"time"
)

type Domain struct {
	Id                uint
	Method_Payment_Id uint
	User_Id           uint
	Total_Qty         uint
	Total_Price       uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Usecase interface {
	InsertTransaction(ctx context.Context, domain *Domain) (Domain, error)
	GetAllTransaction(ctx context.Context, Method_Payment_Id uint, User_Id uint) ([]Domain, error)
	GetTransactionById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertTransaction(ctx context.Context, domain *Domain) (Domain, error)
	GetAllTransaction(ctx context.Context, Method_Payment_Id uint, User_Id uint) ([]Domain, error)
	GetTransactionById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}