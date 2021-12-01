package mocks

import (
	"context"
	"miniproject/business/descriptions"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertDescription(ctx context.Context, domain descriptions.Domain) (descriptions.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 descriptions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, descriptions.Domain) descriptions.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(descriptions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, descriptions.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllDescription(ctx context.Context, search string) ([]descriptions.Domain, error) {
	ret := _m.Called(ctx, search)

	var r0 []descriptions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) []descriptions.Domain); ok {
		r0 = rf(ctx, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]descriptions.Domain)
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

func (_m *Repository) GetDescriptionById(ctx context.Context, id uint) (descriptions.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 descriptions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) descriptions.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(descriptions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) Update(ctx context.Context, domain descriptions.Domain, id uint) (descriptions.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 descriptions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, descriptions.Domain, uint) descriptions.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(descriptions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, descriptions.Domain, uint) error); ok {
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