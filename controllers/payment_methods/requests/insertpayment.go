package requests

import paymentmethods "miniproject/business/payment_methods"

type InsertPayment_Method struct {
	Type string `json:"type"`
}

func (payment_method *InsertPayment_Method) ToDomain() *paymentmethods.Domain {
	return &paymentmethods.Domain{
		Type: payment_method.Type,
	}
}