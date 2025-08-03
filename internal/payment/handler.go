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

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}

func (handler *Handler) MakeTransaction(c *gin.Context) {

	var dto MakeTransactionDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.service.MakeTransaction(dto.From, dto.To, dto.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (handler *Handler) GetLastTransactions(c *gin.Context) {

	countString := c.Query("count")
	if countString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.CountNotProvidedErr.Error()})
		return
	}

	count, err := strconv.Atoi(countString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": common.CountNaNErr.Error()})
		return
	}

	transactions, err := handler.service.GetLastTransactions(count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func (handler *Handler) GetBalance(c *gin.Context) {

	address := c.Param("address")

	balance, err := handler.service.GetBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
