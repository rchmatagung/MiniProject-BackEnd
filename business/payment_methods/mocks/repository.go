package mocks

import (
	"context"
	paymentmethods "miniproject/business/payment_methods"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertPayment_Method(ctx context.Context, domain paymentmethods.Domain) (paymentmethods.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 paymentmethods.Domain
	if rf, ok := ret.Get(0).(func(context.Context, paymentmethods.Domain) paymentmethods.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(paymentmethods.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, paymentmethods.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllPayment_Method(ctx context.Context, search string) ([]paymentmethods.Domain, error) {
	ret := _m.Called(ctx, search)

	var r0 []paymentmethods.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) []paymentmethods.Domain); ok {
		r0 = rf(ctx, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]paymentmethods.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetPayment_MethodById(ctx context.Context, id uint) (paymentmethods.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 paymentmethods.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) paymentmethods.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(paymentmethods.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) Update(ctx context.Context, domain paymentmethods.Domain, id uint) (paymentmethods.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 paymentmethods.Domain
	if rf, ok := ret.Get(0).(func(context.Context, paymentmethods.Domain, uint) paymentmethods.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(paymentmethods.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, paymentmethods.Domain, uint) error); ok {
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