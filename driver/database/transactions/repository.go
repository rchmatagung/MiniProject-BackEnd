package transactions

import (
	"context"
	"errors"
	"miniproject/business/transactions"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		Conn: conn,
	}
}

func (repo *TransactionRepository) InsertTransactions(ctx context.Context, domain *transactions.Domain) (transactions.Domain, error) {
	transaction := FromDomain(*domain)
	err := repo.Conn.Create(&transaction)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return transaction.ToDomain(), nil
}

func (repo *TransactionRepository) GetAllTransaction(ctx context.Context, Method_Payment_Id uint, User_Id uint) ([]transactions.Domain, error) {
	var data []Transaction
	err := repo.Conn.Find(&data)
	if err.Error != nil {
		return []transactions.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *TransactionRepository) GetTransactionById(ctx context.Context, id uint) (transactions.Domain, error) {
	var transaction Transaction
	result := repo.Conn.Find(&transaction, "id = ?", id)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	return transaction.ToDomain(), nil
}

func (repo *TransactionRepository) Update(ctx context.Context, domain transactions.Domain, id uint) (transactions.Domain, error) {
	data := FromDomain(domain)
	if repo.Conn.Save(&data).Error != nil {
		return transactions.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (repo *TransactionRepository) Delete(ctx context.Context, id uint) error {
	transaction := Transaction{}
	err := repo.Conn.Delete(&transaction, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}
