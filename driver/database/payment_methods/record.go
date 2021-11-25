package paymentmethods

import (
	paymentmethods "miniproject/business/payment_methods"
	"time"

	"gorm.io/gorm"
)

type Payment_Methods struct {
	Id        	uint `gorm:"primaryKey"`
	Type      	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain paymentmethods.Domain) Payment_Methods {
	return Payment_Methods{
		Id:        domain.Id,
		Type:      domain.Type,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (payment_method *Payment_Methods) ToDomain() paymentmethods.Domain {
	return paymentmethods.Domain{
		Id:        payment_method.Id,
		Type:      payment_method.Type,
		CreatedAt: payment_method.CreatedAt,
		UpdatedAt: payment_method.UpdatedAt,
	}
}

func ToListDomain(data []Payment_Methods) []paymentmethods.Domain {
	list := []paymentmethods.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}