package paymentmethods

import (
	"context"
	"errors"
	paymentmethods "miniproject/business/payment_methods"

	"gorm.io/gorm"
)

type Payment_MethodRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *Payment_MethodRepository {
	return &Payment_MethodRepository{
		Conn: conn,
	}
}

func (rep *Payment_MethodRepository) InsertPayment_Method(ctx context.Context, domain paymentmethods.Domain) (paymentmethods.Domain, error) {
	payment_method := FromDomain(domain)
	err := rep.Conn.Create(&payment_method)
	if err.Error != nil {
		return paymentmethods.Domain{}, err.Error
	}
	return payment_method.ToDomain(), nil
}

func (rep *Payment_MethodRepository) GetAllPayment_Method(ctx context.Context, search string) ([]paymentmethods.Domain, error) {
	var data []Payment_Methods
	err := rep.Conn.Find(&data)
	if err.Error != nil {
		return []paymentmethods.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *Payment_MethodRepository) GetPayment_MethodById(ctx context.Context, id uint) (paymentmethods.Domain, error) {
	var payment_method Payment_Methods
	result := rep.Conn.Find(&payment_method, "id = ?", id)
	if result.Error != nil {
		return paymentmethods.Domain{}, result.Error
	}

	return payment_method.ToDomain(), nil
}

func (rep *Payment_MethodRepository) Update(ctx context.Context, domain paymentmethods.Domain, id uint) (paymentmethods.Domain, error) {
	data := FromDomain(domain)
	if rep.Conn.Save(&data).Error != nil {
		return paymentmethods.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (rep *Payment_MethodRepository) Delete(ctx context.Context, id uint) error {
	payment_method := Payment_Methods{}
	err := rep.Conn.Delete(&payment_method, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}