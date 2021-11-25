package transactiondetails

import (
	"context"
	"miniproject/business/books"
	"miniproject/business/transactions"
	"time"
)

type Domain struct {
	Id             uint
	Book_Id        uint
	Transaction_Id uint
	Book           books.Domain
	Transaction    transactions.Domain
	Qty            uint
	Price          uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Usecase interface {
	InsertTransaction_Detail(ctx context.Context, domain *Domain) (Domain, error)
	GetAllTransaction_Detail(ctx context.Context, Book_Id uint, Transaction_Id uint) ([]Domain, error)
	GetTransaction_DetailById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	InsertTransaction_Detail(ctx context.Context, domain *Domain) (Domain, error)
	GetAllTransaction_Detail(ctx context.Context, Book_Id uint, Transaction_Id uint) ([]Domain, error)
	GetTransaction_DetailById(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}