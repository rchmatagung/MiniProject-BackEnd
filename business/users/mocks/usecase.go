package mocks

import (
	"context"
	"miniproject/business/users"

	mock "github.com/stretchr/testify/mock"
)

type UseCase struct {
	mock.Mock
}

func (_m *UseCase) Delete(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0). (func(context.Context, uint) error) ; ok{
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

func (_m *UseCase) GetAllUsers(ctx context.Context) ([]users.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []users.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UseCase) GetUserById(ctx context.Context, id uint) (users.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) users.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UseCase) Login(ctx context.Context, email string, password string) (users.Domain, string, error) {
	ret := _m.Called(ctx, email, password)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) users.Domain); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, string) string); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

func (_m *UseCase) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, users.Domain) users.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, users.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UseCase) Update(ctx context.Context, domain users.Domain, id uint) (users.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, users.Domain, uint) users.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, users.Domain, uint) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
