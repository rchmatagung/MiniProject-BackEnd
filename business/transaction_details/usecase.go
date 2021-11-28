package transactiondetails

import (
	"context"
	"time"
)

type Transaction_DetailUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewUseCase(repo Repository, contextTimeout time.Duration) *Transaction_DetailUsecase {
	return &Transaction_DetailUsecase{
		repo: repo,
		contextTimeout: contextTimeout,
	}
}

func (Usecase *Transaction_DetailUsecase) InsertTransaction_Detail(ctx context.Context, domain *Domain) (Domain, error) {
	transaction_detail, err := Usecase.repo.InsertTransaction_Detail(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return transaction_detail, nil
}

func (Usecase *Transaction_DetailUsecase) GetAllTransaction_Detail(ctx context.Context, Book_Id uint, Transaction_Id uint) ([]Domain, error) {
	transaction_detail, err := Usecase.repo.GetAllTransaction_Detail(ctx, Book_Id, Transaction_Id)
	if err != nil {
		return []Domain{}, err
	}
	return transaction_detail, nil
}

func (Usecase *Transaction_DetailUsecase) GetTransaction_DetailById(ctx context.Context, id uint) (Domain, error) {
	transaction_detail, err := Usecase.repo.GetTransaction_DetailById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if transaction_detail.Id == 0 {
		return Domain{}, err
	}
	return transaction_detail, nil
}

func (Usecase *Transaction_DetailUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	transaction_detail, err := Usecase.repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return transaction_detail, nil
}

func (Usecase *Transaction_DetailUsecase) Delete(ctx context.Context, id uint) error {
	err := Usecase.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}