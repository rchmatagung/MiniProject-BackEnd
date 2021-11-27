package responses

import (
	transactiondetails "miniproject/business/transaction_details"
	"time"
)

type Transaction_detail_response struct {
	Id             uint      `json:"id"`
	Book_Id        uint      `json:"book_id"`
	Transaction_Id uint      `json:"transaction_id"`
	Qty            uint      `json:"qty"`
	Price          uint      `json:"price"`
	CreatedAt      time.Time `json:"createdat "`
	UpdatedAt      time.Time `json:"updateat "`
}

type SearchResponse struct {
	Transaction_detail interface{}
}

func FromDomain(domain transactiondetails.Domain) Transaction_detail_response {
	return Transaction_detail_response{
		Id:             domain.Id,
		Book_Id:        domain.Book_Id,
		Transaction_Id: domain.Transaction_Id,
		Qty:            domain.Qty,
		Price:          domain.Price,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func FromTransaction_DetailsAll(domain []transactiondetails.Domain) []Transaction_detail_response {
	var all []Transaction_detail_response
	for _, v := range domain {
		all = append(all, FromDomain(v))
	}
	return all
}