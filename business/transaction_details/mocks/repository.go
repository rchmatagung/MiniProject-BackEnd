package mocks

import (
	"context"
	transactiondetails "miniproject/business/transaction_details"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertTransaction_Detail(ctx context.Context, domain *transactiondetails.Domain) (transactiondetails.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 transactiondetails.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *transactiondetails.Domain) transactiondetails.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(transactiondetails.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *transactiondetails.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllTransaction_Detail(ctx context.Context, Book_Id uint, Transaction_Id uint) ([]transactiondetails.Domain, error) {
	ret := _m.Called(ctx, Book_Id, Transaction_Id)

	var r0 []transactiondetails.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint) []transactiondetails.Domain); ok {
		r0 = rf(ctx, Book_Id, Transaction_Id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactiondetails.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, uint) error); ok {
		r1 = rf(ctx, Book_Id, Transaction_Id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetTransaction_DetailById(ctx context.Context, id uint) (transactiondetails.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 transactiondetails.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) transactiondetails.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(transactiondetails.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) Update(ctx context.Context, domain transactiondetails.Domain, id uint) (transactiondetails.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 transactiondetails.Domain
	if rf, ok := ret.Get(0).(func(context.Context, transactiondetails.Domain, uint) transactiondetails.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(transactiondetails.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, transactiondetails.Domain, uint) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) Delete(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}