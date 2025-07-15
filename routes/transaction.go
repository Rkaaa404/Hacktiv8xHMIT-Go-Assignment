package routes

import (
	"net/http"
	"assignment/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterTransactionRoutes(r *gin.Engine, transactions *[]models.Transaction, products *[]models.Product) {
	r.GET("/transactions", getTransactions(transactions))
	r.GET("/transactions/:id", getTransactionByID(transactions))
	r.POST("/transactions", createTransaction(transactions, products))
}

func getTransactions(transactions *[]models.Transaction) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.StandardResponse{Message: "Transactions fetched successfully", Data: *transactions})
	}
}

func getTransactionByID(transactions *[]models.Transaction) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		for _, t := range *transactions {
			if t.ID == id {
				c.JSON(http.StatusOK, models.StandardResponse{Message: "Transaction fetched successfully", Data: t})
				return
			}
		}
		c.JSON(http.StatusNotFound, models.StandardResponse{Error: "Transaction not found"})
	}
}

func createTransaction(transactions *[]models.Transaction, products *[]models.Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.TransactionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.StandardResponse{Error: err.Error()})
			return
		}

		productIndex := -1
		var product models.Product
		for i, p := range *products {
			if p.ID == req.ProductID {
				product = p
				productIndex = i
				break
			}
		}

		if productIndex == -1 {
			c.JSON(http.StatusNotFound, models.StandardResponse{Error: "Product not found"})
			return
		}

		if product.Stock < req.Quantity {
			c.JSON(http.StatusBadRequest, models.StandardResponse{Error: "Insufficient stock"})
			return
		}

		(*products)[productIndex].Stock -= req.Quantity

		newTransaction := models.Transaction{
			ID:        uuid.New().String(),
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
			Total:     product.Price * float64(req.Quantity),
		}
		*transactions = append(*transactions, newTransaction)

		c.JSON(http.StatusCreated, models.StandardResponse{Message: "Transaction success", Data: newTransaction})
	}
}