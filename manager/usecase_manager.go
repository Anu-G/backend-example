package manager

import "wmb-rest-api/usecase"

type UseCaseManagerInterface interface {
	CustomerUseCase() usecase.CustomerUseCaseInterface
	DiscountUseCase() usecase.DiscountUseCaseInterface
	TableUseCase() usecase.TableUseCaseInterface
	MenuUseCase() usecase.MenuUseCaseInterface
	TrxUseCase() usecase.TrxUseCaseInterface
}

type useCaseManager struct {
	repo RepositoryManagerInterface
}

func NewUseCase(manager RepositoryManagerInterface) UseCaseManagerInterface {
	return &useCaseManager{
		repo: manager,
	}
}

func (um *useCaseManager) CustomerUseCase() usecase.CustomerUseCaseInterface {
	return usecase.NewCustomerUseCase(um.repo.CustomerRepo(), um.DiscountUseCase())
}

func (um *useCaseManager) DiscountUseCase() usecase.DiscountUseCaseInterface {
	return usecase.NewDiscountUseCase(um.repo.DiscountRepo())
}

func (um *useCaseManager) TableUseCase() usecase.TableUseCaseInterface {
	return usecase.NewTableUseCase(um.repo.TableRepo())
}

func (um *useCaseManager) MenuUseCase() usecase.MenuUseCaseInterface {
	return usecase.NewMenuUseCase(um.repo.MenuRepo())
}

func (um *useCaseManager) TrxUseCase() usecase.TrxUseCaseInterface {
	return usecase.NewTrxUseCase(um.repo.BillRepo(), um.repo.TrxTypeRepo(), um.repo.LopeiRepo(),
		um.CustomerUseCase(), um.TableUseCase(), um.MenuUseCase(), um.DiscountUseCase())
}
