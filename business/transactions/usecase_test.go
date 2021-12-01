package transactions_test

import (
	"context"
	"errors"
	"miniproject/business/transactions"
	"miniproject/business/transactions/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transactionRepository = mocks.Repository{Mock: mock.Mock{}}
var transactionService transactions.Usecase
var transactionDomain transactions.Domain
var listTransactionDomain []transactions.Domain

func setup() {
	transactionService = transactions.NewUseCase(&transactionRepository, time.Hour*10)
	transactionDomain = transactions.Domain{
		Id:                1,
		User_Id:           1,
		Method_Payment_Id: 1,
		Total_Qty:         1,
		Total_Price:       100000,
	}
	listTransactionDomain = append(listTransactionDomain, transactionDomain)
}

func TestInsertTransaction(t *testing.T) {
	setup()
	transactionRepository.On("InsertTransaction", mock.Anything, mock.Anything).Return(transactionDomain, nil)
	t.Run("Test Case 1 | Success Insert Transaction", func(t *testing.T) {
		transaction, err := transactionService.InsertTransaction(context.Background(), &transactions.Domain{})

		assert.NoError(t, err)
		assert.Equal(t, transactionDomain, transaction)
	})
}

func TestGetListTransactions(t *testing.T) {
	t.Run("Test case 1 | Success GetListTransactions", func(t *testing.T) {
		setup()
		transactionRepository.On("GetAllTransaction", mock.Anything, mock.Anything).Return(listTransactionDomain, nil).Once()
		data, err := transactionService.GetAllTransaction(context.Background(), transactionDomain.Method_Payment_Id, transactionDomain.User_Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listTransactionDomain))
	})

	t.Run("Test case 2 | Error GetListTransactions", func(t *testing.T) {
		setup()
		transactionRepository.On("GetAllTransaction", mock.Anything, mock.Anything).Return([]transactions.Domain{}, errors.New("Transactions Not Found")).Once()
		data, err := transactionService.GetAllTransaction(context.Background(), 0, 0)

		assert.Error(t, err)
		assert.Equal(t, data, []transactions.Domain{})
	})
}

func TestSearchTransactionById(t *testing.T) {
	t.Run("Test case 1 | Success GetListTransactionById", func(t *testing.T) {
		setup()
		transactionRepository.On("GetTransactionById", mock.Anything, mock.AnythingOfType("uint")).Return(transactionDomain, nil).Once()
		data, err := transactionService.GetTransactionById(context.Background(), transactionDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetListTransactionById(transaction Id = 0)", func(t *testing.T) {
		setup()
		transactionDomain.Id = 0
		transactionRepository.On("GetTransactionById", mock.Anything, mock.AnythingOfType("uint")).Return(transactionDomain, nil).Once()
		data, err := transactionService.GetTransactionById(context.Background(), transactionDomain.Id)

		assert.NoError(t, err)
		assert.Equal(t, data, transactions.Domain{})
	})

	t.Run("Test case 3 | Error GetListTransactionById", func(t *testing.T) {
		setup()
		transactionRepository.On("GetTransactionById", mock.Anything, mock.AnythingOfType("uint")).Return(transactions.Domain{}, nil).Once()
		data, err := transactionService.GetTransactionById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, transactions.Domain{})
	})
}

func TestUpdateTransaction(t *testing.T) {
	t.Run("Test case 1 | Success Update Transaction", func(t *testing.T) {
		setup()
		transactionRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(transactionDomain, nil).Once()
		data, err := transactionService.Update(context.Background(), transactionDomain, transactionDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Transaction", func(t *testing.T) {
		setup()
		transactionRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(transactionDomain, errors.New("transactions Not Found")).Once()
		data, err := transactionService.Update(context.Background(), transactionDomain, transactionDomain.Id)

		assert.Equal(t, data, transactions.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteTransaction(t *testing.T) {
	t.Run("Test case 1 | Success Delete Transaction", func(t *testing.T) {
		setup()
		transactionRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := transactionService.Delete(context.Background(), transactionDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete transaction", func(t *testing.T) {
		setup()
		transactionRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Transactions  Not Found")).Once()
		err := transactionService.Delete(context.Background(), transactionDomain.Id)

		assert.NotEqual(t, err, errors.New("Transactions Not Found"))
		assert.Error(t, err)
	})
}