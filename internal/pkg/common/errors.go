package common

import "errors"

type HttpError struct {
	Error string `json:"error"`
}

var (
	CountNotProvidedErr  = errors.New("argument 'count' was not provided")
	CountNaNErr          = errors.New("argument 'count' is not a number")
	InsufficientFundsErr = errors.New("wallet with provided 'from' address does not have enough funds")
	WalletNotExistsErr   = errors.New("wallet with provided %s address does not exist")
)
