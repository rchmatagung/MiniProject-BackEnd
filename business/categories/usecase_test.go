package categories_test

import (
	"context"
	"errors"
	"miniproject/business/categories"
	"miniproject/business/categories/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = mocks.Repository{Mock: mock.Mock{}}
var categoryService categories.UseCase
var categoryDomain categories.Domain
var AllcategoryDomain []categories.Domain

func setup() {
	categoryService = categories.NewCategoryUseCase(&categoryRepository, time.Hour*10)
	categoryDomain = categories.Domain{
		Id:       1,
		Category: "Teknik Informatika",
	}
	AllcategoryDomain = append(AllcategoryDomain, categoryDomain)
}

func TestInsertCategory(t *testing.T) {
	setup()
	categoryRepository.On("InsertCategory", mock.Anything, mock.Anything).Return(categoryDomain, nil)
	t.Run("Test Case 1 | Success Insert Category", func(t *testing.T) {
		category, err := categoryService.InsertCategory(context.Background(), categories.Domain{
			Id:       1,
			Category: "Teknik Informatika",
		})

		assert.NoError(t, err)
		assert.Equal(t, categoryDomain, category)
	})
}

func TestGetAllCategory(t *testing.T) {
	t.Run("Test case 1 | Success SearchCategories", func(t *testing.T) {
		setup()
		categoryRepository.On("GetAllCategory", mock.Anything, mock.Anything).Return(AllcategoryDomain, nil).Once()
		data, err := categoryService.GetAllCategory(context.Background(), categoryDomain.Category)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(AllcategoryDomain))
	})

	t.Run("Test case 2 | Error SearchCategories(search empty)", func(t *testing.T) {
		setup()
		categoryRepository.On("GetAllCategory", mock.Anything, mock.Anything).Return([]categories.Domain{}, errors.New("Categories Not Found")).Once()
		data, err := categoryService.GetAllCategory(context.Background(), "")

		assert.Error(t, err)
		assert.Equal(t, data, []categories.Domain{})
	})
}

func TestGetCategoryById(t *testing.T) {
	t.Run("Test case 1 | Success GetCategoryById", func(t *testing.T) {
		setup()
		categoryRepository.On("GetCategoryById", mock.Anything, mock.AnythingOfType("uint")).Return(categoryDomain, nil).Once()
		data, err := categoryService.GetCategoryById(context.Background(), categoryDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetCategoryById(category Id = 0)", func(t *testing.T) {
		setup()
		categoryDomain.Id = 0
		categoryRepository.On("GetCategoryById", mock.Anything, mock.AnythingOfType("uint")).Return(categoryDomain, nil).Once()
		data, err := categoryService.GetCategoryById(context.Background(), categoryDomain.Id)

		assert.NoError(t, err)
		assert.Equal(t, data, categories.Domain{})
	})

	t.Run("Test case 3 | Error SearchCategoryById", func(t *testing.T) {
		setup()
		categoryRepository.On("GetCategoryById", mock.Anything, mock.AnythingOfType("uint")).Return(categories.Domain{}, nil).Once()
		data, err := categoryService.GetCategoryById(context.Background(), 7)

		assert.NoError(t, err)
		assert.Equal(t, data, categories.Domain{})
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("Test case 1 | Success Update Category", func(t *testing.T) {
		setup()
		categoryRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(categoryDomain, nil).Once()
		data, err := categoryService.Update(context.Background(), categoryDomain, categoryDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update Category", func(t *testing.T) {
		setup()
		categoryRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(categoryDomain, errors.New("Categories Not Found")).Once()
		data, err := categoryService.Update(context.Background(), categoryDomain, categoryDomain.Id)

		assert.Equal(t, data, categories.Domain{})
		assert.Error(t, err)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("Test case 1 | Success Delete Category", func(t *testing.T) {
		setup()
		categoryRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := categoryService.Delete(context.Background(), categoryDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete Category", func(t *testing.T) {
		setup()
		categoryRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Categories  Not Found")).Once()
		err := categoryService.Delete(context.Background(), categoryDomain.Id)

		assert.NotEqual(t, err, errors.New("Categories Not Found"))
		assert.Error(t, err)
	})
}