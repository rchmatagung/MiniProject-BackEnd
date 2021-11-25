package books

import (
	"context"
	"time"
)

type BookUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewUseCase(repo Repository, contextTimeout time.Duration) *BookUsecase {
	return &BookUsecase{
		repo: repo,
		contextTimeout: contextTimeout,
	}
}

func (UseCase *BookUsecase) InsertBook(ctx context.Context, domain *Domain) (Domain, error) {
	book, err := UseCase.repo.InsertBook(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return book, nil
}

func (UseCase *BookUsecase) GetAllBook(ctx context.Context, search string) ([]Domain, error) {
	book, err := UseCase.repo.GetAllBook(ctx, search)
	if err != nil {
		return []Domain{}, err
	}

	return book, nil
}

func (UseCase *BookUsecase) GetBookById(ctx context.Context, id uint) (Domain, error) {
	book, err := UseCase.repo.GetBookById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if book.Id == 0 {
		return Domain{}, err
	}
	return book, nil
}

func (UseCase *BookUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	book, err := UseCase.repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return book, nil
}

func (UseCase *BookUsecase) Delete(ctx context.Context, id uint) error {
	err := UseCase.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}