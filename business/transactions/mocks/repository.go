package mocks

import (
	"context"
	"miniproject/business/transactions"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertTransaction(ctx context.Context, domain *transactions.Domain) (transactions.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *transactions.Domain) transactions.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *transactions.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllTransaction(ctx context.Context, Method_Payment_Id uint, User_Id uint) ([]transactions.Domain, error) {
	ret := _m.Called(ctx, Method_Payment_Id, User_Id)

	var r0 []transactions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint) []transactions.Domain); ok {
		r0 = rf(ctx, Method_Payment_Id, User_Id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, uint) error); ok {
		r1 = rf(ctx, Method_Payment_Id, User_Id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetTransactionById(ctx context.Context, id uint) (transactions.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) transactions.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) Update(ctx context.Context, domain transactions.Domain, id uint) (transactions.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, transactions.Domain, uint) transactions.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, transactions.Domain, uint) error); ok {
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