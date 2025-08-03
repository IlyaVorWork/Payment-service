package common

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HttpError - структура для возврата ошибки в ответе HTTP
type HttpError struct {
	Error string `json:"error"`
}

var (
	CountNotProvidedErr  = errors.New("argument 'count' was not provided")
	CountNaNErr          = errors.New("argument 'count' is not a number")
	InsufficientFundsErr = errors.New("wallet with provided 'from' address does not have enough funds")
	WalletNotExistsErr   = errors.New("wallet with provided %s address does not exist")
)

func NewErrorHandler() gin.HandlerFunc {
	codeMap := getErrorCodesMap()

	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0].Error()

		code, ok := codeMap[err]
		if !ok {
			code = http.StatusInternalServerError
		}

		jsonErrs := HttpError{Error: err}

		c.JSON(code, jsonErrs)
	}
}

func getErrorCodesMap() map[string]int {
	codeMap := make(map[string]int)
	codeMap[CountNotProvidedErr.Error()] = http.StatusBadRequest
	codeMap[CountNaNErr.Error()] = http.StatusBadRequest
	codeMap[InsufficientFundsErr.Error()] = http.StatusBadRequest
	codeMap[fmt.Sprintf(WalletNotExistsErr.Error(), "'from'")] = http.StatusBadRequest
	codeMap[fmt.Sprintf(WalletNotExistsErr.Error(), "'to'")] = http.StatusBadRequest
	codeMap[fmt.Sprintf(WalletNotExistsErr.Error(), "")] = http.StatusBadRequest

	return codeMap
}
