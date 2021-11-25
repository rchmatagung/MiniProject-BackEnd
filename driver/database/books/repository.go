package books

import (
	"context"
	"errors"
	"miniproject/business/books"

	"gorm.io/gorm"
)

type BookRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *BookRepository {
	return &BookRepository{
		Conn: conn,
	}
}

func (repo *BookRepository) InsertBook(ctx context.Context, domain *books.Domain) (books.Domain, error) {
	book := FromDomain(*domain)
	err := repo.Conn.Create(&book)
	if err.Error != nil {
		return books.Domain{}, err.Error
	}
	return book.ToDomain(), nil
}

func (repo *BookRepository) GetAllBook(ctx context.Context, search string) ([]books.Domain, error) {
	var data []Books
	err := repo.Conn.Find(&data)
	if err.Error != nil {
		return []books.Domain{}, err.Error
	}
	return AllBook(data), nil
}

func (repo *BookRepository) GetBookById(ctx context.Context, id uint) (books.Domain, error) {
	var book Books
	err := repo.Conn.Find(&book, "id = ?", id)
	if err.Error != nil {
		return books.Domain{}, err.Error
	}
	return book.ToDomain(), nil
}

func (repo *BookRepository) Update(ctx context.Context, domain books.Domain, id uint) (books.Domain, error) {
	data := FromDomain(domain)
	if repo.Conn.Find(&data).Error != nil {
		return books.Domain{}, errors.New("Bad Request")
	}
	return data.ToDomain(), nil
}	

func (repo *BookRepository) Delete(ctx context.Context, id uint) error {
	book := Books{}
	err := repo.Conn.Delete(&book, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("Id Not Found")
	}
	return nil
}



