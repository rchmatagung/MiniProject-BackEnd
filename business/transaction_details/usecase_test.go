package transactiondetails_test

import (
	"context"
	"errors"
	transactiondetails "miniproject/business/transaction_details"
	"miniproject/business/transaction_details/mocks"
	"miniproject/business/transactions"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transaction_detailRepository = mocks.Repository{Mock: mock.Mock{}}
var transaction_detailService transactiondetails.Usecase
var transaction_detailDomain transactiondetails.Domain
var listTransaction_DetailDomain []transactiondetails.Domain

func setup() {
	transaction_detailService = transactiondetails.NewUseCase(&transaction_detailRepository, time.Hour*10)
	transaction_detailDomain = transactiondetails.Domain{
		Id:             1,
		Book_Id:        1,
		Transaction_Id: 1,
		Qty:            1,
		Price:          100000,
	}
	listTransaction_DetailDomain = append(listTransaction_DetailDomain, transaction_detailDomain)
}

func TestInsertTransaction_Detail(t *testing.T) {
	setup()
	transaction_detailRepository.On("InsertTransaction_Detail", mock.Anything, mock.Anything).Return(transaction_detailDomain, nil)
	t.Run("Test Case 1 | Success Insert Transaction_Detail", func(t *testing.T) {
		transaction_detail, err := transaction_detailService.InsertTransaction_Detail(context.Background(), &transactiondetails.Domain{})

		assert.NoError(t, err)
		assert.Equal(t, transaction_detailDomain, transaction_detail)
	})
}

func TestGetListTransaction_Details(t *testing.T) {
	t.Run("Test case 1 | Success GetListTransaction_Details", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("GetAllTransaction_Detail", mock.Anything, mock.Anything, mock.Anything).Return(listTransaction_DetailDomain, nil).Once()
		data, err := transaction_detailService.GetAllTransaction_Detail(context.Background(), transaction_detailDomain.Book_Id, transaction_detailDomain.Transaction_Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listTransaction_DetailDomain))
	})

	t.Run("Test case 2 | Error GetListTransaction_Details", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("GetAllTransaction_Detail", mock.Anything, mock.Anything, mock.Anything).Return([]transactiondetails.Domain{}, errors.New("Transaction_Details Not Found")).Once()
		data, err := transaction_detailService.GetAllTransaction_Detail(context.Background(), 0, 0)

		assert.Error(t, err)
		assert.Equal(t, data, []transactiondetails.Domain{})
	})
}

func TestSearchTransaction_DetailById(t *testing.T) {
	t.Run("Test case 1 | Success GetListTransaction_DetailById", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("GetTransaction_DetailById", mock.Anything, mock.AnythingOfType("uint")).Return(transaction_detailDomain, nil).Once()
		data, err := transaction_detailService.GetTransaction_DetailById(context.Background(), transaction_detailDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetListTransaction_DetailById(transaction Id = 0)", func(t *testing.T) {
		setup()
		transaction_detailDomain.Id = 0
		transaction_detailRepository.On("GetTransaction_DetailById", mock.Anything, mock.AnythingOfType("uint")).Return(transaction_detailDomain, nil).Once()
		data, err := transaction_detailService.GetTransaction_DetailById(context.Background(), transaction_detailDomain.Id)

		assert.NoError(t, err)
		assert.NotEqual(t, data, transactions.Domain{})
	})

	t.Run("Test case 3 | Error GetListTransaction_DetailById", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("GetTransaction_DetailById", mock.Anything, mock.AnythingOfType("uint")).Return(transactiondetails.Domain{}, nil).Once()
		data, err := transaction_detailService.GetTransaction_DetailById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, transactiondetails.Domain{})
	})
}

func TestUpdateTransaction_Detail(t *testing.T) {
	t.Run("Test case 1 | Success Update Transaction_Detail", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(transaction_detailDomain, nil).Once()
		data, err := transaction_detailService.Update(context.Background(), transaction_detailDomain, transaction_detailDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Transaction_Detail", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(transaction_detailDomain, errors.New("transactiondetails Not Found")).Once()
		data, err := transaction_detailService.Update(context.Background(), transaction_detailDomain, transaction_detailDomain.Id)

		assert.NotEqual(t, data, transactions.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteTransaction_Detail(t *testing.T) {
	t.Run("Test case 1 | Success Delete Transaction_Detail", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := transaction_detailService.Delete(context.Background(), transaction_detailDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete transaction_detail", func(t *testing.T) {
		setup()
		transaction_detailRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Transaction_details  Not Found")).Once()
		err := transaction_detailService.Delete(context.Background(), transaction_detailDomain.Id)

		assert.NotEqual(t, err, errors.New("Transaction_details Not Found"))
		assert.Error(t, err)
	})
}