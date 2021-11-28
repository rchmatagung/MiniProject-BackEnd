package descriptions

import (
	"context"
	"errors"
	"miniproject/business/descriptions"

	"gorm.io/gorm"
)

type DescriptionRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *DescriptionRepository {
	return &DescriptionRepository{
		Conn: conn,
	}
}

func (repo *DescriptionRepository) InsertDescription(ctx context.Context, domain descriptions.Domain) (descriptions.Domain, error) {
	category := FromDomain(domain)
	err := repo.Conn.Create(&category)
	if err.Error != nil {
		return descriptions.Domain{}, err.Error
	}
	return category.ToDomain(), nil
}

func (repo *DescriptionRepository) GetAllDescription(ctx context.Context, search string) ([]descriptions.Domain, error) {
	var data []Descriptions
	err := repo.Conn.Find(&data)
	if err.Error != nil {
		return []descriptions.Domain{}, err.Error
	}
	return AllDescription(data), nil
}

func (repo *DescriptionRepository) GetDescriptionById(ctx context.Context, id uint) (descriptions.Domain, error) {
	var description Descriptions
	err := repo.Conn.Find(&description, "id = ?", id)
	if err.Error != nil {
		return descriptions.Domain{}, err.Error
	}
	return description.ToDomain(), nil
}

func (repo *DescriptionRepository) Update(ctx context.Context, domain descriptions.Domain, id uint) (descriptions.Domain, error) {
	data := FromDomain(domain)
	if repo.Conn.Save(&data).Error != nil {
		return descriptions.Domain{}, errors.New("Bad Request")
	}
	return data.ToDomain(), nil
}

func (repo *DescriptionRepository) Delete(ctx context.Context, id uint) error {
	description := Descriptions{}
	err := repo.Conn.Delete(&description, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("Id Not Found")
	}
	return nil
}