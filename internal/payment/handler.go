package payment

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment_service/internal/pkg/common"
	"strconv"
)

type IService interface {
	MakeTransaction(from, to string, amount float64) error
	GetLastTransactions(count int) ([]Transaction, error)
	GetBalance(address string) (float64, error)
}

// Handler представляет собой слой обработчиков HTTP-запросов
type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}

// MakeTransaction [POST] /api/send
// Получает в теле запроса данные MakeTransactionDTO
// Ответы:
// 200 - В случае успеха
// 400 - В случае предоставления неверных входных данных или недостатка баланса
// 500 - В случае ошибки со стороны сервера
func (handler *Handler) MakeTransaction(c *gin.Context) {

	var dto MakeTransactionDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&dto); err != nil {
		_ = c.Error(err)
		return
	}

	err := handler.service.MakeTransaction(dto.From, dto.To, dto.Amount)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

// GetLastTransactions [GET] /api/transactions?count=N
// Получает в параметрах запроса count - число записей, которое необходимо вернуть
// Ответы:
// 200 - В случае успеха, возвращает данные GetLastTransactionsOut
// 400 - В случае отсутствия параметра count или его неверного формата
// 500 - В случае ошибки со стороны сервера
func (handler *Handler) GetLastTransactions(c *gin.Context) {

	countString := c.Query("count")
	if countString == "" {
		_ = c.Error(common.CountNotProvidedErr)
		return
	}

	count, err := strconv.Atoi(countString)
	if err != nil {
		_ = c.Error(common.CountNaNErr)
		return
	}

	transactions, err := handler.service.GetLastTransactions(count)
	if err != nil {
		_ = c.Error(err)
		return
	}

	out := GetLastTransactionsOut{
		Transactions: transactions,
	}
	c.JSON(http.StatusOK, out)
}

// GetBalance [GET] /api/wallet/:address/balance
// Получает в адресе запроса address - адрес кошелька, баланс которого необходимо вернуть
// Ответы:
// 200 - В случае успеха, возвращает данные GetBalanceOut
// 400 - В случае отсутствия параметра count или его неверного формата
// 500 - В случае ошибки со стороны сервера
func (handler *Handler) GetBalance(c *gin.Context) {

	address := c.Param("address")

	balance, err := handler.service.GetBalance(address)
	if err != nil {
		_ = c.Error(err)
		return
	}

	out := GetBalanceOut{
		Balance: balance,
	}
	c.JSON(http.StatusOK, out)
}
