package paymentmethods_test

import (
	"context"
	"errors"
	paymentmethods "miniproject/business/payment_methods"
	"miniproject/business/payment_methods/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var payment_methodRepository = mocks.Repository{Mock: mock.Mock{}}
var payment_methodService paymentmethods.Usecase
var payment_methodDomain paymentmethods.Domain
var AllPayment_MethodDomain []paymentmethods.Domain

func setup() {
	payment_methodService = paymentmethods.NewUseCase(&payment_methodRepository, time.Hour*10)
	payment_methodDomain = paymentmethods.Domain{
		Id:   1,
		Type: "Bank BRI: 0098 0114 9875 503",
	}
	AllPayment_MethodDomain = append(AllPayment_MethodDomain, payment_methodDomain)
}

func TestInsertPayment_Method(t *testing.T) {
	setup()
	payment_methodRepository.On("InsertPayment_Method", mock.Anything, mock.Anything).Return(payment_methodDomain, nil)
	t.Run("Test Case 1 | Success Insert Payment_Method", func(t *testing.T) {
		payment_method, err := payment_methodService.InsertPayment_Method(context.Background(), paymentmethods.Domain{
			Id:   1,
			Type: "ank BRI: 0098 0114 9875 503",
		})

		assert.NoError(t, err)
		assert.Equal(t, payment_methodDomain, payment_method)
	})

	setup()
	payment_methodRepository.On("InsertPayment_Method", mock.Anything, mock.Anything).Return(payment_methodDomain, errors.New("method payment empty")).Once()
	t.Run("Test Case 2 | Error Insert Payment_Method", func(t *testing.T) {
		payment_method, err := payment_methodService.InsertPayment_Method(context.Background(), paymentmethods.Domain{
			Id:   1,
			Type: "",
		})

		assert.Error(t, err)
		assert.NotNil(t, payment_method)
	})
}

func TestGetAllPayment_Method(t *testing.T) {
	t.Run("Test case 1 | Success SearchPayment_Methods", func(t *testing.T) {
		setup()
		payment_methodRepository.On("GetAllPayment_Method", mock.Anything, mock.Anything).Return(AllPayment_MethodDomain, nil).Once()
		data, err := payment_methodService.GetAllPayment_Method(context.Background(), payment_methodDomain.Type)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(AllPayment_MethodDomain))
	})

	t.Run("Test case 2 | Error SearchPayment_Methods(search empty)", func(t *testing.T) {
		setup()
		payment_methodRepository.On("GetAllPayment_Method", mock.Anything, mock.Anything).Return([]paymentmethods.Domain{}, errors.New("Payment_Methods Not Found")).Once()
		data, err := payment_methodService.GetAllPayment_Method(context.Background(), "")

		assert.Error(t, err)
		assert.Equal(t, data, []paymentmethods.Domain{})
	})
}

func TestGetPayment_MethodById(t *testing.T) {
	t.Run("Test case 1 | Success SearchPayment_MethodById", func(t *testing.T) {
		setup()
		payment_methodRepository.On("GetPayment_MethodById", mock.Anything, mock.AnythingOfType("uint")).Return(payment_methodDomain, nil).Once()
		data, err := payment_methodService.GetPayment_MethodById(context.Background(), payment_methodDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error SearchPayment_MethodById(payment_method Id = 0)", func(t *testing.T) {
		setup()
		payment_methodDomain.Id = 0
		payment_methodRepository.On("GetPayment_MethodById", mock.Anything, mock.AnythingOfType("uint")).Return(payment_methodDomain, nil).Once()
		data, err := payment_methodService.GetPayment_MethodById(context.Background(), payment_methodDomain.Id)

		assert.NoError(t, err)
		assert.Equal(t, data, paymentmethods.Domain{})
	})

	t.Run("Test case 3 | Error SearchPayment_MethodById", func(t *testing.T) {
		setup()
		payment_methodRepository.On("GetPayment_MethodById", mock.Anything, mock.AnythingOfType("uint")).Return(paymentmethods.Domain{}, nil).Once()
		data, err := payment_methodService.GetPayment_MethodById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, paymentmethods.Domain{})
	})
}

func TestUpdatePayment_Method(t *testing.T) {
	t.Run("Test case 1 | Success Update Payment_Method", func(t *testing.T) {
		setup()
		payment_methodRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(payment_methodDomain, nil).Once()
		data, err := payment_methodService.Update(context.Background(), payment_methodDomain, payment_methodDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Payment_Method", func(t *testing.T) {
		setup()
		payment_methodRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(payment_methodDomain, errors.New("Payment_Methods Not Found")).Once()
		data, err := payment_methodService.Update(context.Background(), payment_methodDomain, payment_methodDomain.Id)

		assert.Equal(t, data, paymentmethods.Domain{})
		assert.Error(t, err)
	})
}

func TestDeletePayment_Method(t *testing.T) {
	t.Run("Test case 1 | Success Delete Payment_Method", func(t *testing.T) {
		setup()
		payment_methodRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := payment_methodService.Delete(context.Background(), payment_methodDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete Payment_Method", func(t *testing.T) {
		setup()
		payment_methodRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Payment_Methods  Not Found")).Once()
		err := payment_methodService.Delete(context.Background(), payment_methodDomain.Id)

		assert.NotEqual(t, err, errors.New("Payment_Methods Not Found"))
		assert.Error(t, err)
	})
}

