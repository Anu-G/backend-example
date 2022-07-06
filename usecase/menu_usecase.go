package usecase

import (
	"wmb-rest-api/model/dto"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/repository"
)

type MenuUseCaseInterface interface {
	GetMenu(t *entity.Menu) error
	GetMenuPrice(t *entity.Menu) (foundPrice entity.MenuPrice, err error)
	GetMenuPriceById(t *entity.MenuPrice) error
	FindMenuPriceAndMenu(t *entity.MenuPrice) (foundMenu entity.Menu, err error)
	UpdateMenu(mr *dto.MenuRequest) (updatedMenu entity.Menu, err error)
	DeleteMenu(m *entity.Menu) error
	CreateMenu(mr *dto.MenuRequest) (createdMenu entity.Menu, err error)
}

type menuUseCase struct {
	repo repository.MenuRepositoryInterface
}

func NewMenuUseCase(repo repository.MenuRepositoryInterface) MenuUseCaseInterface {
	return &menuUseCase{
		repo: repo,
	}
}

func (mu *menuUseCase) GetMenu(t *entity.Menu) error {
	if err := mu.repo.FindById(t); err != nil {
		return err
	}

	if menuPrice, err := mu.repo.FindLatestMenuPrice(t); err != nil {
		return err
	} else {
		t.MenuPrices = nil
		t.MenuPrices = append(t.MenuPrices, menuPrice)
	}
	return nil
}

func (mu *menuUseCase) GetMenuPrice(t *entity.Menu) (foundPrice entity.MenuPrice, err error) {
	if foundPrice, err = mu.repo.FindLatestMenuPrice(t); err != nil {
		return foundPrice, err
	}
	return foundPrice, nil
}

func (mu *menuUseCase) GetMenuPriceById(t *entity.MenuPrice) error {
	return mu.repo.FindMenuPriceById(t)
}

func (mu *menuUseCase) FindMenuPriceAndMenu(t *entity.MenuPrice) (foundMenu entity.Menu, err error) {
	return mu.repo.FindMenuByMenuPrice(t)
}

func (mu *menuUseCase) UpdateMenu(mr *dto.MenuRequest) (updatedMenu entity.Menu, err error) {
	updatedMenu.ID = mr.MenuID
	updateMenuPrice := entity.MenuPrice{Price: mr.MenuPrice}
	if err = mu.repo.FindByIdAndPrice(&updatedMenu, &updateMenuPrice, "MenuPrices"); err != nil {
		return updatedMenu, err
	}

	if mr.MenuName != "" {
		updatedMenu.MenuName = mr.MenuName
	}
	updatedMenu.MenuPrices = append(updatedMenu.MenuPrices, updateMenuPrice)

	if err = mu.repo.Update(&updatedMenu); err != nil {
		return updatedMenu, err
	}

	return updatedMenu, err
}

func (mu *menuUseCase) DeleteMenu(m *entity.Menu) error {
	if err := mu.repo.FindByIdPreload(m, "MenuPrices"); err != nil {
		return err
	}

	if err := mu.repo.DeleteAssociation(m, "MenuPrices", m.MenuPrices); err != nil {
		return err
	}

	if err := mu.repo.Delete(m); err != nil {
		return err
	}
	return nil
}

func (mu *menuUseCase) CreateMenu(mr *dto.MenuRequest) (createdMenu entity.Menu, err error) {
	createdMenu.MenuName = mr.MenuName
	createdMenu.MenuPrices = append(createdMenu.MenuPrices, entity.MenuPrice{Price: mr.MenuPrice})

	if err = mu.repo.Create(&createdMenu); err != nil {
		return createdMenu, err
	}
	return createdMenu, err
}
