package usecase

import (
	"errors"

	"wmb-rest-api/model/dto"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/repository"

	"gorm.io/gorm"
)

type CustomerUseCaseInterface interface {
	ReadOrCreateCustomer(c *entity.Customer) error
	FindById(c *entity.Customer) error
	FindByPhone(c *entity.Customer) error
	FindByIdPreload(c *entity.Customer) error
	UpdateCustomer(cr *dto.CustomerRequest) (updatedCustomer entity.Customer, err error)
	DeleteCustomer(c *entity.Customer) error
}

type customerUseCase struct {
	repo       repository.CustomerRepositoryInterface
	discountUC DiscountUseCaseInterface
}

func NewCustomerUseCase(repo repository.CustomerRepositoryInterface, du DiscountUseCaseInterface) CustomerUseCaseInterface {
	return &customerUseCase{
		repo:       repo,
		discountUC: du,
	}
}

func (cu *customerUseCase) ReadOrCreateCustomer(c *entity.Customer) error {
	foundCust, err := cu.repo.FindFirtst(map[string]interface{}{"mobile_phone_no": c.MobilePhoneNo})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = cu.repo.Create(c); err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}

	if foundCust.ID != 0 {
		return errors.New("phone number already registered, please login with your username and password")
	}
	return nil
}

func (cu *customerUseCase) FindById(c *entity.Customer) error {
	return cu.repo.FindById(c)
}

func (cu *customerUseCase) FindByPhone(c *entity.Customer) error {
	return cu.repo.FindByPhone(c)
}

func (cu *customerUseCase) FindByIdPreload(c *entity.Customer) error {
	return cu.repo.FindByIdPreload(c, "Discounts")
}

func (cu *customerUseCase) UpdateCustomer(cr *dto.CustomerRequest) (updatedCustomer entity.Customer, err error) {
	var discountFound entity.Discount

	updatedCustomer.ID = cr.CustomerID
	if err = cu.repo.FindByIdPreload(&updatedCustomer, "Discounts"); err != nil {
		return updatedCustomer, err
	}

	if cr.CustomerName != "" {
		updatedCustomer.CustomerName = cr.CustomerName
	}

	if cr.MobilePhoneNo != "" {
		updatedCustomer.MobilePhoneNo = cr.MobilePhoneNo
	}
	updatedCustomer.IsMember = cr.IsMember

	if cr.DiscountID != 0 {
		discountFound.ID = cr.DiscountID
		if err = cu.discountUC.GetDiscountByID(&discountFound); err != nil {
			return updatedCustomer, err
		}

		updatedCustomer.Discounts = append(updatedCustomer.Discounts, &discountFound)
	}
	if err = cu.repo.Update(&updatedCustomer); err != nil {
		return updatedCustomer, err
	}
	return updatedCustomer, err
}

func (cu *customerUseCase) DeleteCustomer(c *entity.Customer) error {
	if err := cu.repo.FindByIdPreload(c, "Discounts"); err != nil {
		return err
	}

	if err := cu.repo.DeleteAssociation(c, "Discounts", c.Discounts); err != nil {
		return err
	}

	if err := cu.repo.Delete(c); err != nil {
		return err
	}
	return nil
}
