package api

import (
	"main/db"
	"main/intercepter"
	"main/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupTransactionAPI -
func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.GET("/transaction", getTransaction)
		transactionAPI.POST("/transaction", intercepter.JwtVerify, createtransaction)
	}
}

// TransactionResult - transaction api
type TransactionResult struct {
	ID            uint
	Total         float64
	Paid          float64
	change        float64
	PaymentType   string
	PaymentDetail string
	OrderList     string
	Staff         string
	CreatedAt     time.Time
}

// Transaction - get transaction api
// func getTransaction(c *gin.Context) {
// 	var transaction []model.Transaction
// 	db.GetDB().Find(&transaction)
// 	c.JSON(http.StatusOK, transaction)
// }

// func getTransaction(c *gin.Context) {
// 	var transaction []model.Transaction
// 	db.GetDB().Debug().Raw("SELECT transactions.id, total, paid, change, payment_type, payment_details, order_list, users.username as Staff, transactions.created_at FROM transactions join users on transactions.staff_id = users.id", nil).Scan(&transaction)
// 	c.JSON(http.StatusOK, transaction)
// }

func getTransaction(c *gin.Context) {
	var result []TransactionResult
	db.GetDB().Debug().Raw("SELECT transactions.id, total, paid, change, payment_type, payment_detail, order_list, users.username as Staff, transactions.created_at FROM transactions left join users on transactions.staff_id = users.id", nil).Scan(&result)
	c.JSON(200, result)
}

// Createtranaction - createtransaction api
func createtransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBind(&transaction); err == nil {
		transaction.StaffID = c.GetString("jwt_staff_id")
		transaction.CreatedAt = time.Now()
		db.GetDB().Create(&transaction)
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": transaction})
	} else {
		c.JSON(404, gin.H{"result": "nok"})
	}
}
