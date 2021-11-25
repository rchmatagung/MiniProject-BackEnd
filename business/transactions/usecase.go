package transactions

import (
	"context"
	"time"
)

type TransactionUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewUseCase(repo Repository, contextTimeout time.Duration) *TransactionUsecase {
	return &TransactionUsecase{
		repo:           repo,
		contextTimeout: contextTimeout,
	}
}

func (UseCase *TransactionUsecase) InsertTransaction(ctx context.Context, domain *Domain) (Domain, error) {
	transaction, err := UseCase.repo.InsertTransaction(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (UseCase *TransactionUsecase) GetAllTransaction(ctx context.Context, Method_Payment_Id uint, User_Id uint) ([]Domain, error) {
	transaction, err := UseCase.repo.GetAllTransaction(ctx, 0, 0)
	if err != nil {
		return []Domain{}, err
	}

	return transaction, nil
}

func (UseCase *TransactionUsecase) GetTransactionById(ctx context.Context, id uint) (Domain, error) {
	transaction, err := UseCase.repo.GetTransactionById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if transaction.Id == 0 {
		return Domain{}, err
	}
	return transaction, nil
}

func (UseCase *TransactionUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	transaction, err := UseCase.repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}

func (UseCase *TransactionUsecase) Delete(ctx context.Context, id uint) error {
	err := UseCase.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}