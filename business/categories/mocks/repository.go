package mocks

import (
	"context"
	"miniproject/business/categories"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) InsertCategory(ctx context.Context, domain categories.Domain) (categories.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 categories.Domain
	if rf, ok := ret.Get(0).(func(context.Context, categories.Domain) categories.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(categories.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, categories.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllCategory(ctx context.Context, search string) ([]categories.Domain, error) {
	ret := _m.Called(ctx, search)

	var r0 []categories.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) []categories.Domain); ok {
		r0 = rf(ctx, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]categories.Domain)
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

func (_m *Repository) GetCategoryById(ctx context.Context, id uint) (categories.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 categories.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) categories.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(categories.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) Update(ctx context.Context, domain categories.Domain, id uint) (categories.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 categories.Domain
	if rf, ok := ret.Get(0).(func(context.Context, categories.Domain, uint) categories.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(categories.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, categories.Domain, uint) error); ok {
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
