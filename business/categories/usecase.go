package categories

import (
	"context"
	"errors"
	"time"
)

type CategoryUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewUseCase(repo Repository, contextTimeout time.Duration) *CategoryUsecase {
	return &CategoryUsecase{
		repo:           repo,
		contextTimeout: contextTimeout,
	}
}

func (UseCase *CategoryUsecase) InsertCategory(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Category == "" {
		return Domain{}, errors.New("category empty")
	}

	category, err := UseCase.repo.InsertCategory(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return category, nil
}

func (UseCase *CategoryUsecase) GetListCategory(ctx context.Context, search string) ([]Domain, error) {
	category, err := UseCase.repo.GetAllCategory(ctx, search)
	if err != nil {
		return []Domain{}, err
	}

	return category, nil
}

func (UseCase *CategoryUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	category, err := UseCase.repo.GetCategoryById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if category.Id == 0 {
		return Domain{}, err
	}
	return category, nil
}

func (UseCase *CategoryUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	category, err := UseCase.repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return category, nil
}

func (UseCase *CategoryUsecase) Delete(ctx context.Context, id uint) error {
	err := UseCase.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}