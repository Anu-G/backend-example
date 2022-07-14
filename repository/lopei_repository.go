package repository

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"wmb-rest-api/service"
)

type LopeiRepositoryInterface interface {
	CheckBalance(phoneNumber string) (float64, error)
	DoPayment(phoneNumber string, amount float64) error
}

type lopeiRepository struct {
	client service.LopeiPaymentClient
}

func NewLopeiRepo(client service.LopeiPaymentClient) LopeiRepositoryInterface {
	return &lopeiRepository{
		client: client,
	}
}

func (bpr *lopeiRepository) CheckBalance(phoneNumber string) (float64, error) {
	res, err := bpr.client.CheckBalance(context.Background(), &service.CheckBalanceMessage{
		LopeiId: phoneNumber,
	})
	if err != nil {
		return 0, err
	}

	split := strings.Split(res.Result, ":")
	balStr := split[2]
	regex, err := regexp.Compile(`\d+`)
	if err != nil {
		return 0, err
	}
	balStr = regex.FindString(balStr)

	balance, err := strconv.ParseFloat(balStr, 64)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (bpr *lopeiRepository) DoPayment(phoneNumber string, amount float64) error {
	if res, err := bpr.client.DoPayment(context.Background(), &service.PaymentMessage{
		LopeiId: phoneNumber,
		Amount:  amount,
	}); err != nil {
		return err
	} else if res.Error != nil {
		return errors.New(res.Error.Message)
	}
	return nil
}
