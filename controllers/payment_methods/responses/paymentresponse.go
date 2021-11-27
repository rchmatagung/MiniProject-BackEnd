package responses

import (
	paymentmethods "miniproject/business/payment_methods"
	"time"
)

type Payment_MethodResponse struct {
	Id        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdat "`
	UpdatedAt time.Time `json:"updateat "`
}

type SearchResponse struct {
	Payment_Method interface{}
}

func FromDomain(domain paymentmethods.Domain) Payment_MethodResponse {
	return Payment_MethodResponse{
		Id:        domain.Id,
		Type:      domain.Type,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromPayment_MethodsAll(domain []paymentmethods.Domain) []Payment_MethodResponse {
	var all []Payment_MethodResponse
	for _, v := range domain {
		all = append(all, FromDomain(v))
	}
	return all
}