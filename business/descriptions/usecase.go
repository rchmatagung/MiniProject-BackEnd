package descriptions

import (
	"context"
	"errors"
	"time"
)

type DescriptionUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewUseCase(repo Repository, contextTimeout time.Duration) *DescriptionUsecase {
	return &DescriptionUsecase{
		repo:           repo,
		contextTimeout: contextTimeout,
	}
}

func (UseCase *DescriptionUsecase) InsertDescription(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Description == "" {
		return Domain{}, errors.New("Description empty")
	}

	description, err := UseCase.repo.InsertDescription(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return description, nil
}

func (UseCase *DescriptionUsecase) GetAllDescription(ctx context.Context, search string) ([]Domain, error) {
	description, err := UseCase.repo.GetAllDescription(ctx, search)
	if err != nil {
		return []Domain{}, err
	}

	return description, nil
}

func (UseCase *DescriptionUsecase) GetDescriptionById(ctx context.Context, id uint) (Domain, error) {
	description, err := UseCase.repo.GetDescriptionById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if description.Id == 0 {
		return Domain{}, err
	}
	return description, nil
}

func (UseCase *DescriptionUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	description, err := UseCase.repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return description, nil
}

func (UseCase *DescriptionUsecase) Delete(ctx context.Context, id uint) error {
	err := UseCase.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}