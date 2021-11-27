package requests

import "miniproject/business/transactions"

type Transaction_Search struct {
	Method_Payment_Id uint `json:"method_payment_Id"`
	User_Id           uint `json:"user_id"`
	Total_Qty         uint `json:"total_qty"`
	Total_Price       uint `json:"total_price"`
}

func ToDomain(search Transaction_Search) transactions.Domain {
	return transactions.Domain{
		Method_Payment_Id: search.Method_Payment_Id,
		User_Id:           search.User_Id,
		Total_Qty:         search.Total_Qty,
		Total_Price:       search.Total_Price,
	}
}