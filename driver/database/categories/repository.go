package categories

import (
	"context"
	"errors"
	"miniproject/business/categories"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	Conn *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		Conn: conn,
	}
}

func (repo *CategoryRepository) InsertCategory(ctx context.Context, domain categories.Domain) (categories.Domain, error) {
	category := FromDomain(domain)
	err := repo.Conn.Create(&category)
	if err.Error != nil {
		return categories.Domain{}, err.Error
	}
	return category.ToDomain(), nil
}

func (repo *CategoryRepository) GetAllCategory(ctx context.Context, search string) ([]categories.Domain, error) {
	var data []Categories
	err := repo.Conn.Find(&data)
	if err.Error != nil {
		return []categories.Domain{}, err.Error
	}
	return AllCategory(data), nil
}

func (repo *CategoryRepository) GetCategoryById(ctx context.Context, id uint) (categories.Domain, error) {
	var category Categories
	err := repo.Conn.Find(&category, "id = ?", id)
	if err.Error != nil {
		return categories.Domain{}, err.Error
	}
	return category.ToDomain(), nil
}

func (repo *CategoryRepository) Update(ctx context.Context, domain categories.Domain, id uint) (categories.Domain, error) {
	data := FromDomain(domain)
	if repo.Conn.Find(&data).Error != nil {
		return categories.Domain{}, errors.New("Bad Request")
	}
	return data.ToDomain(), nil
}	

func (repo *CategoryRepository) Delete(ctx context.Context, id uint) error {
	category := Categories{}
	err := repo.Conn.Delete(&category, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("Id Not Found")
	}
	return nil
}