package repository

import (
	"database/sql"
	"errors"
	"time"

	"wmb-rest-api/model/entity"

	"gorm.io/gorm"
)

type BillRepositoryInterface interface {
	CreateTransaction(c *entity.Customer, t *entity.Table, tt *entity.TransactionType, details *[]entity.BillDetail) (int, error)
	CreateBillPayment(bp *entity.BillPayment) error
	FindById(b *entity.Bill) error
	FindAllBillDetail(by map[string]interface{}) ([]entity.BillDetail, error)
	FindAllByDate(date string) ([]entity.Bill, error)
	Delete(b *entity.Bill) error
}

type billRepository struct {
	db *gorm.DB
}

func NewBillRepo(db *gorm.DB) BillRepositoryInterface {
	return &billRepository{
		db: db,
	}
}

func (br *billRepository) CreateTransaction(c *entity.Customer, t *entity.Table, tt *entity.TransactionType, details *[]entity.BillDetail) (int, error) {
	var (
		bill     entity.Bill
		err      error
		discount entity.Discount
	)

	bill.TransactionDate = time.Now()
	bill.CustomerID = c.ID
	bill.TransactionTypeID = tt.ID
	bill.TableID = sql.NullInt64{Int64: int64(t.ID)}
	if t.ID != 0 {
		bill.TableID.Valid = true
	}

	// Select One Customer Discount
	for _, data := range c.Discounts {
		if data.Pct > discount.Pct {
			discount = *data
		}
	}
	bill.DiscountID = sql.NullInt64{Int64: int64(discount.ID)}
	if discount.ID != 0 {
		bill.DiscountID.Valid = true
	}

	startTrx := br.db.Begin()
	startTrx.Transaction(
		func(tx *gorm.DB) error {
			if err = tx.Create(&bill).Error; err != nil {
				return err
			}

			for _, data := range *details {
				data.BillID = bill.ID
				if err = tx.Create(&data).Error; err != nil {
					return err
				}
			}

			if t.ID != 0 {
				if err = tx.Model(&t).Update("is_available", false).Error; err != nil {
					return err
				}
			}

			tx.Commit()
			return nil
		}, &sql.TxOptions{ReadOnly: true},
	)
	return int(bill.ID), err
}

func (br *billRepository) FindById(b *entity.Bill) error {
	return br.db.First(&b).Error
}

func (br *billRepository) CreateBillPayment(bp *entity.BillPayment) error {
	return br.db.Create(&bp).Error
}

func (br *billRepository) FindAllBillDetail(by map[string]interface{}) ([]entity.BillDetail, error) {
	var details []entity.BillDetail
	res := br.db.Where(by).Find(&details)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return details, nil
		} else {
			return details, err
		}
	}
	return details, nil
}

func (br *billRepository) FindAllByDate(date string) ([]entity.Bill, error) {
	var bills []entity.Bill
	if err := br.db.Raw("SELECT * FROM t_bill WHERE date(transaction_date) = ? AND t_bill.deleted_at IS NULL", date).Scan(&bills).Error; err != nil {
		return bills, err
	}
	return bills, nil
}

func (br *billRepository) Delete(b *entity.Bill) error {
	return br.db.Delete(b).Error
}
