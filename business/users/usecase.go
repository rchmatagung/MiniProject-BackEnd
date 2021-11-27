package users

import (
	"context"
	"errors"
	"log"
	"miniproject/app/middleware"
	"time"
)

type UserUsecase struct {
	repo Repository
	contextTimeout time.Duration
	jwtAuth	*middleware.ConfigJWT
}

func NewUseCase(UserRepo Repository, contextTimeout time.Duration, jwtAuth *middleware.ConfigJWT) UseCase {
	return &UserUsecase{
		repo: UserRepo,
		contextTimeout: contextTimeout,
		jwtAuth: jwtAuth,
	}
}

func (UseCase *UserUsecase) Login(ctx context.Context, email string, password string) (Domain, string, error) {
	if email == "" {
		return Domain{}, "", errors.New("Email empty")
	}

	if password == "" {
		return Domain{}, "", errors.New("Password empty")
	}

	user, err := UseCase.repo.GetByEmail(ctx, email)

	if err != nil {
		return Domain{}, "", err
	}

	token, errToken := UseCase.jwtAuth.GenerateTokenJWT(user.Id)
	if errToken != nil {
		log.Println(errToken)
	}
	if token == "" {
		return Domain{}, "", errors.New("authentication failed: invalid user credentials")
	}
	return user,token, nil
}

func (UseCase *UserUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("Email Empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("Password Empty")
	}

	data, err := UseCase.repo.GetByEmail(ctx, domain.Email)

	if data.Id > 0 {
		return Domain{}, errors.New("Email Already Used")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("Password Required")
	}

	user, err := UseCase.repo.Register(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (UseCase *UserUsecase) GetAllUsers(ctx context.Context) ([]Domain, error) {
	user, err := UseCase.repo.GetAllUsers(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, err
}

func (UseCase *UserUsecase) GetUserById(ctx context.Context, Id uint) (Domain, error) {
	user, err := UseCase.repo.GetUserById(ctx, Id)
	if err != nil {
		return Domain{}, err
	}
	if user.Id == 0 {
		return Domain{}, err
	}
	return user, nil
}

func (UseCase *UserUsecase) Update(ctx context.Context, domain Domain, Id uint) (Domain, error) {
	domain.Id = (Id)
	user, err := UseCase.repo.Update(ctx, domain, Id)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (UseCase *UserUsecase) Delete(ctx context.Context, Id uint) error {
	err := UseCase.repo.Delete(ctx, Id)
	if err != nil {
		return err
	}
	return nil
}




