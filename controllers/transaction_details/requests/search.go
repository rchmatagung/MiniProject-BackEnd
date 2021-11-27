package requests

import (
	"miniproject/business/books"
	transactiondetails "miniproject/business/transaction_details"
	"miniproject/business/transactions"
)

type Transaction_Detail_Search struct {
	Book        books.Domain        `json:"book"`
	Transaction transactions.Domain `json:"transaction"`
	Qty         uint                `json:"qty"`
	Price       uint                `json:"price"`
}

func ToDomain(search Transaction_Detail_Search) transactiondetails.Domain {
	return transactiondetails.Domain{
		Book:        books.Domain{},
		Transaction: transactions.Domain{},
		Qty:         search.Qty,
		Price:       search.Price,
	}
}