package manager

import "wmb-rest-api/repository"

type RepositoryManagerInterface interface {
	CustomerRepo() repository.CustomerRepositoryInterface
	MenuRepo() repository.MenuRepositoryInterface
	TableRepo() repository.TableRepositoryInterface
	DiscountRepo() repository.DiscountRepositoryInterface
	TrxTypeRepo() repository.TrxTypeInterface
	BillRepo() repository.BillRepositoryInterface
	LopeiRepo() repository.LopeiRepositoryInterface
	AuthRepo() repository.AuthRepositoryInterface
}

type repositoryManager struct {
	dbCon InfraManagerInterface
}

func NewRepo(dbCon InfraManagerInterface) RepositoryManagerInterface {
	return &repositoryManager{
		dbCon: dbCon,
	}
}

func (rm *repositoryManager) CustomerRepo() repository.CustomerRepositoryInterface {
	return repository.NewCustomerRepo(rm.dbCon.DBCon())
}

func (rm *repositoryManager) MenuRepo() repository.MenuRepositoryInterface {
	return repository.NewMenuRepo(rm.dbCon.DBCon())
}

func (rm *repositoryManager) TableRepo() repository.TableRepositoryInterface {
	return repository.NewTableRepo(rm.dbCon.DBCon())
}

func (rm *repositoryManager) DiscountRepo() repository.DiscountRepositoryInterface {
	return repository.NewDiscountRepo(rm.dbCon.DBCon())
}

func (rm *repositoryManager) TrxTypeRepo() repository.TrxTypeInterface {
	return repository.NewTrxTypeRepo(rm.dbCon.DBCon())
}

func (rm *repositoryManager) BillRepo() repository.BillRepositoryInterface {
	return repository.NewBillRepo(rm.dbCon.DBCon())
}

func (rm *repositoryManager) LopeiRepo() repository.LopeiRepositoryInterface {
	return repository.NewLopeiRepo(rm.dbCon.LopeiCon())
}

func (rm *repositoryManager) AuthRepo() repository.AuthRepositoryInterface {
	return repository.NewAuthRepo(rm.dbCon.DBCon())
}
