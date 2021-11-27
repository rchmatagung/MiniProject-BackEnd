package requests

import paymentmethods "miniproject/business/payment_methods"

type Payment_MethodSearch struct {
	Type string `json:"type"`
}

func ToDomain(search Payment_MethodSearch) paymentmethods.Domain {
	return paymentmethods.Domain{
		Type: search.Type,
	}
}