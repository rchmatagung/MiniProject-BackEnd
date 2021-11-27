package requests

import transactiondetails "miniproject/business/transaction_details"

type InsertTransaction_detail struct {
	Book_Id        uint `json:"book_id"`
	Transaction_Id uint `json:"transaction_id"`
	Qty            uint `json:"qty"`
	Price          uint `json:"price"`
}

func (transaction_detail *InsertTransaction_detail) ToDomain() *transactiondetails.Domain {
	return &transactiondetails.Domain{
		Book_Id:        transaction_detail.Book_Id,
		Transaction_Id: transaction_detail.Transaction_Id,
		Qty:            transaction_detail.Qty,
		Price:          transaction_detail.Price,
	}
}
