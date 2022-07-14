package usecase

import (
	"wmb-rest-api/model/entity"
	"wmb-rest-api/repository"
)

type AuthUseCaseInterface interface {
	CreateUser(uc *entity.UserCredential) error
	FindUser(uc *entity.UserCredential) error
}

type authUseCase struct {
	repo repository.AuthRepositoryInterface
}

func NewAuthUseCase(repo repository.AuthRepositoryInterface) AuthUseCaseInterface {
	return &authUseCase{
		repo: repo,
	}
}

func (au *authUseCase) CreateUser(uc *entity.UserCredential) error {
	return au.repo.CreateUser(uc)
}

func (au *authUseCase) FindUser(uc *entity.UserCredential) error {
	return au.repo.FindUser(uc)
}
