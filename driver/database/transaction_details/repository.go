package transactiondetails

import (
	"context"
	"errors"
	transactiondetails "miniproject/business/transaction_details"

	"gorm.io/gorm"
)

type Transaction_DetailRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *Transaction_DetailRepository {
	return &Transaction_DetailRepository{
		Conn: conn,
	}
}


func (repo *Transaction_DetailRepository) InsertTransaction_Detail(ctx context.Context, domain *transactiondetails.Domain) (transactiondetails.Domain, error) {
	transaction_detail := FromDomain(*domain)
	err := repo.Conn.Create(&transaction_detail)
	if err.Error != nil {
		return transactiondetails.Domain{}, err.Error
	}
	return transaction_detail.ToDomain(), nil
}

func (repo *Transaction_DetailRepository) GetAllTransaction_Detail(ctx context.Context, Book_Id uint, Transaction_Id uint) ([]transactiondetails.Domain, error) {
	var data []Transaction_Detail
	err := repo.Conn.Find(&data)
	if err.Error != nil {
		return []transactiondetails.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *Transaction_DetailRepository) GetTransaction_DetailById(ctx context.Context, id uint) (transactiondetails.Domain, error) {
	var transaction_detail Transaction_Detail
	result := repo.Conn.Find(&transaction_detail, "id = ?", id)
	if result.Error != nil {
		return transactiondetails.Domain{}, result.Error
	}

	return transaction_detail.ToDomain(), nil
}


func (repo *Transaction_DetailRepository) Update(ctx context.Context, domain transactiondetails.Domain, id uint) (transactiondetails.Domain, error) {
	data := FromDomain(domain)
	if repo.Conn.Save(&data).Error != nil {
		return transactiondetails.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (repo *Transaction_DetailRepository) Delete(ctx context.Context, id uint) error {
	transaction_detail := Transaction_Detail{}
	err := repo.Conn.Delete(&transaction_detail, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}