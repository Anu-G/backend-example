package usecase

import (
	"errors"
	"strconv"

	"wmb-rest-api/model/dto"
	"wmb-rest-api/model/entity"
	"wmb-rest-api/repository"

	"gorm.io/gorm"
)

type TrxUseCaseInterface interface {
	CreateTransaction(crtrx *dto.CreateTransaction) (int, error)
	PayAndFinishTransaction(pay *dto.PaymentMethod) (dto.BillPrintOut, error)
	GetRevenue(rev *dto.Revenue) error
	CheckBalance(c *entity.Customer) (dto.LopeiBalance, error)
}

type trxUseCase struct {
	repo        repository.BillRepositoryInterface
	trxTypeRepo repository.TrxTypeInterface
	lopeiRepo   repository.LopeiRepositoryInterface
	customerUC  CustomerUseCaseInterface
	tableUC     TableUseCaseInterface
	menuUC      MenuUseCaseInterface
	discountUC  DiscountUseCaseInterface
}

func NewTrxUseCase(repo repository.BillRepositoryInterface, ttp repository.TrxTypeInterface, bp repository.LopeiRepositoryInterface,
	cu CustomerUseCaseInterface, tu TableUseCaseInterface, mu MenuUseCaseInterface, du DiscountUseCaseInterface,
) TrxUseCaseInterface {
	return &trxUseCase{
		repo:        repo,
		trxTypeRepo: ttp,
		lopeiRepo:   bp,
		customerUC:  cu,
		tableUC:     tu,
		menuUC:      mu,
		discountUC:  du,
	}
}

func (tu *trxUseCase) CreateTransaction(crtrx *dto.CreateTransaction) (int, error) {
	var (
		err         error
		billID      int
		customer    entity.Customer
		table       entity.Table
		trxType     entity.TransactionType
		menu        entity.Menu
		order       entity.MenuPrice
		billDetails []entity.BillDetail
	)

	// Customer Validation
	customer = crtrx.Customer
	if err = tu.customerUC.FindByPhone(&customer); err != nil {
		return billID, err
	}

	// Transaction Type Validation
	trxType.ID = crtrx.TransactionTypeID
	if err = tu.trxTypeRepo.FindById(&trxType); err != nil {
		return billID, err
	}

	// Table Validation
	table.ID = crtrx.TableID
	if trxType.ID == "TA" && table.ID != 0 {
		notAllowed := errors.New("take away not allowed to choose table")
		return billID, notAllowed
	} else if trxType.ID == "DI" && table.ID != 0 {
		if err = tu.tableUC.GetTable(&table); err != nil {
			return billID, err
		} else if !table.IsAvailable {
			tableNotAvailable := errors.New("table not available, please select another")
			return billID, tableNotAvailable
		}
	} else if trxType.ID == "DI" && table.ID == 0 {
		tableRequired := errors.New("please choose a table for dine in")
		return billID, tableRequired
	}

	// Menu and Menu Price Validation
	for _, data := range crtrx.OrderMenus {
		menu.ID = data.MenuID
		if err = tu.menuUC.GetMenu(&menu); err != nil {
			return billID, err
		}

		if order, err = tu.menuUC.GetMenuPrice(&menu); err != nil {
			return billID, err
		}

		billDetails = append(billDetails, entity.BillDetail{
			MenuPriceID: order.ID,
			Qty:         data.Qty,
		})
	}

	// Create Bill
	if billID, err = tu.repo.CreateTransaction(&customer, &table, &trxType, &billDetails); err != nil {
		return billID, err
	}
	return billID, nil
}

func (tu *trxUseCase) PayAndFinishTransaction(pay *dto.PaymentMethod) (dto.BillPrintOut, error) {
	var (
		err             error
		customer        entity.Customer
		bill            entity.Bill
		billPayment     entity.BillPayment
		transactionType entity.TransactionType
		discount        entity.Discount
		menu            entity.Menu
		menuPrice       entity.MenuPrice
		orders          []entity.BillDetail
		printOut        dto.BillPrintOut
	)

	// Validate Bill
	bill.ID = pay.BillId
	if err = tu.repo.FindById(&bill); err != nil {
		return printOut, err
	}
	customer.ID = bill.CustomerID
	transactionType.ID = bill.TransactionTypeID

	// Get Customer Data
	if err = tu.customerUC.FindById(&customer); err != nil {
		return printOut, err
	}

	// Get Transaction Type
	if err = tu.trxTypeRepo.FindById(&transactionType); err != nil {
		return printOut, err
	}

	// Get Discount Percent
	if bill.DiscountID.Int64 != 0 {
		discount.ID = bill.Discount.ID
		if err = tu.discountUC.GetDiscountByID(&discount); err != nil {
			return printOut, err
		}
	}

	// Get All Orders
	if orders, err = tu.repo.FindAllBillDetail(map[string]interface{}{"bill_id": bill.ID}); err != nil {
		return printOut, err
	}

	for _, data := range orders {
		menuPrice.ID = data.MenuPriceID
		if menu, err = tu.menuUC.FindMenuPriceAndMenu(&menuPrice); err != nil {
			return dto.BillPrintOut{}, err
		}

		summary := dto.HistoryMenuOrder{
			MenuName: menu.MenuName,
			Price:    menuPrice.Price,
			Qty:      data.Qty,
			Subtotal: menuPrice.Price * float64(data.Qty),
		}

		printOut.Orders = append(printOut.Orders, summary)
		printOut.GrandTotal += summary.Subtotal
	}

	printOut.Discount = discount.Pct
	if discount.Pct != 0 {
		discNum := (float64(100-discount.Pct) / 100)
		printOut.GrandTotal = printOut.GrandTotal * discNum
	}

	// pay
	if pay.PaymentMethod != "lopei" && pay.PaymentMethod != "cash" {
		return dto.BillPrintOut{}, errors.New("payment method invalid")
	}

	if pay.PaymentMethod == "lopei" {
		if err = tu.lopeiRepo.DoPayment(customer.MobilePhoneNo, printOut.GrandTotal); err != nil {
			return dto.BillPrintOut{}, err
		}
	}
	billPayment.BillID = bill.ID
	billPayment.PaymentMethod = pay.PaymentMethod
	billPayment.TotalPayment = printOut.GrandTotal
	if err = tu.repo.CreateBillPayment(&billPayment); err != nil {
		return dto.BillPrintOut{}, err
	}

	printOut.BillID = bill.ID
	printOut.TransactionDate = bill.TransactionDate.Format("2 Jan 2006 15:04:05")
	printOut.CustomerName = customer.CustomerName
	if bill.TransactionTypeID != "TA" {
		printOut.TransactionType = transactionType.Description
		printOut.Table = strconv.FormatInt(bill.TableID.Int64, 10)
		tu.tableUC.UpdateTableAvailability(&entity.Table{Model: gorm.Model{ID: uint(bill.TableID.Int64)}}, true)
	} else {
		printOut.TransactionType = transactionType.Description
	}

	return printOut, err
}

func (tu *trxUseCase) GetRevenue(rev *dto.Revenue) error {
	var (
		bills       []entity.Bill
		billDetails []entity.BillDetail
		menuPrice   entity.MenuPrice
		discount    entity.Discount
		err         error
	)
	if bills, err = tu.repo.FindAllByDate(rev.TransactionDate); err != nil {
		return err
	}

	for _, data := range bills {
		var subtotal float64

		if billDetails, err = tu.repo.FindAllBillDetail(map[string]interface{}{"bill_id": data.ID}); err != nil {
			return err
		}

		for _, details := range billDetails {
			menuPrice.ID = details.MenuPriceID
			if err = tu.menuUC.GetMenuPriceById(&menuPrice); err != nil {
				return err
			}
			subtotal += menuPrice.Price * float64(details.Qty)
		}

		if data.DiscountID.Int64 != 0 {
			discount.ID = data.Discount.ID
			if err = tu.discountUC.GetDiscountByID(&discount); err != nil {
				return err
			}
			discNum := (float64(100-discount.Pct) / 100)
			subtotal = subtotal * discNum
		}

		rev.TotalRevenue += subtotal
	}
	return err
}

func (tu *trxUseCase) CheckBalance(c *entity.Customer) (dto.LopeiBalance, error) {
	var (
		lopeiData dto.LopeiBalance
		balance   float64
		err       error
	)

	if err = tu.customerUC.FindByPhone(c); err != nil {
		return lopeiData, err
	}
	balance, err = tu.lopeiRepo.CheckBalance(c.MobilePhoneNo)
	if err != nil {
		return lopeiData, err
	}
	lopeiData.CustomerName = c.CustomerName
	lopeiData.MobilePhoneNo = c.MobilePhoneNo
	lopeiData.Balance = balance

	return lopeiData, nil
}
