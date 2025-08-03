package payment

import (
	"errors"
	"fmt"
	"math"
	"payment_service/internal/pkg/common"
)

type IProvider interface {
	MakeTransaction(from, to string, amount float64) error
	GetLastTransactions(count int) ([]Transaction, error)
	GetBalance(address string) (float64, error)
}

type Service struct {
	provider IProvider
}

func NewService(provider IProvider) *Service {
	return &Service{
		provider: provider,
	}
}

func (service *Service) MakeTransaction(from, to string, amount float64) error {

	roundedAmount := math.Floor(amount*100) / 100

	fromBalance, err := service.provider.GetBalance(from)
	if err != nil {
		return errors.New(fmt.Sprintf(common.WalletNotExistsErr.Error(), "'from'"))
	}
	if fromBalance < roundedAmount {
		return common.InsufficientFundsErr
	}

	_, err = service.provider.GetBalance(to)
	if err != nil {
		return errors.New(fmt.Sprintf(common.WalletNotExistsErr.Error(), "'to'"))
	}

	err = service.provider.MakeTransaction(from, to, roundedAmount)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) GetLastTransactions(count int) ([]Transaction, error) {

	transactions, err := service.provider.GetLastTransactions(count)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (service *Service) GetBalance(address string) (float64, error) {

	balance, err := service.provider.GetBalance(address)
	if err != nil {
		return 0, errors.New(fmt.Sprintf(common.WalletNotExistsErr.Error(), ""))
	}

	return balance, nil
}
