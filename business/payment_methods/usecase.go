package paymentmethods

import (
	"context"
	"errors"
	"time"
)

type Payment_MethodUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewUseCase(repo Repository, contextTimeout time.Duration) *Payment_MethodUsecase {
	return &Payment_MethodUsecase{
		repo:           repo,
		contextTimeout: contextTimeout,
	}
}

func (Usecase *Payment_MethodUsecase) InsertPayment_Method(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Type == "" {
		return Domain{}, errors.New("method payment empty")
	}

	payment_method, err := Usecase.repo.InsertPayment_Method(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return payment_method, nil
}

func (Usecase *Payment_MethodUsecase) GetAllPayment_Method(ctx context.Context, search string) ([]Domain, error) {
	payment_method, err := Usecase.repo.GetAllPayment_Method(ctx, search)
	if err != nil {
		return []Domain{}, err
	}
	return payment_method, nil
}

func (Usecase *Payment_MethodUsecase) GetPayment_MethodById(ctx context.Context, id uint) (Domain, error) {
	payment_method, err := Usecase.repo.GetPayment_MethodById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if payment_method.Id == 0 {
		return Domain{}, err
	}
	return payment_method, nil
}

func (Usecase *Payment_MethodUsecase) Update(ctx context.Context, domain Domain, id uint) (Domain, error) {
	domain.Id = (id)
	payment_method, err := Usecase.repo.Update(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}
	return payment_method, nil
}

func (Usecase *Payment_MethodUsecase) Delete(ctx context.Context, id uint) error {
	err := Usecase.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}