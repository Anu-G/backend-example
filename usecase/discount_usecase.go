package usecase

import (
	"wmb-rest-api/model/entity"
	"wmb-rest-api/repository"
)

type DiscountUseCaseInterface interface {
	GetDiscountByID(d *entity.Discount) error
	UpdateDiscount(d *entity.Discount) error
	DeleteDiscount(d *entity.Discount) error
	CreateDiscount(d *entity.Discount) error
}

type discountUseCase struct {
	repo repository.DiscountRepositoryInterface
}

func NewDiscountUseCase(repo repository.DiscountRepositoryInterface) DiscountUseCaseInterface {
	return &discountUseCase{
		repo: repo,
	}
}

func (du *discountUseCase) GetDiscountByID(d *entity.Discount) error {
	return du.repo.FindById(d)
}

func (du *discountUseCase) UpdateDiscount(d *entity.Discount) error {
	return du.repo.Update(d)
}

func (du *discountUseCase) DeleteDiscount(d *entity.Discount) error {
	if err := du.repo.FindByIdPreload(d, "Customers"); err != nil {
		return err
	}

	if err := du.repo.DeleteAssociation(d, "Customers", d.Customers); err != nil {
		return err
	}

	if err := du.repo.Delete(d); err != nil {
		return err
	}
	return nil
}

func (du *discountUseCase) CreateDiscount(d *entity.Discount) error {
	return du.repo.Create(d)
}
