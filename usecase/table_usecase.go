package usecase

import (
	"wmb-rest-api/model/entity"
	"wmb-rest-api/repository"
)

type TableUseCaseInterface interface {
	GetTable(t *entity.Table) error
	UpdateTable(t *entity.Table) error
	UpdateTableAvailability(t *entity.Table, isAvailable bool) error
	DeleteTable(t *entity.Table) error
	CreateTable(t *entity.Table) error
}

type tableUseCase struct {
	repo repository.TableRepositoryInterface
}

func NewTableUseCase(repo repository.TableRepositoryInterface) TableUseCaseInterface {
	return &tableUseCase{
		repo: repo,
	}
}

func (tu *tableUseCase) GetTable(t *entity.Table) error {
	if err := tu.repo.FindById(t); err != nil {
		return err
	}
	return nil
}

func (tu *tableUseCase) UpdateTable(t *entity.Table) error {
	return tu.repo.Update(t)
}

func (tu *tableUseCase) UpdateTableAvailability(t *entity.Table, isAvailable bool) error {
	t.IsAvailable = isAvailable
	return tu.repo.Update(t)
}

func (tu *tableUseCase) DeleteTable(t *entity.Table) error {
	return tu.repo.Delete(t)
}

func (tu *tableUseCase) CreateTable(t *entity.Table) error {
	return tu.repo.Create(t)
}
