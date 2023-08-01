package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swapnika/jwt-auth/middleware"
	"github.com/swapnika/jwt-auth/transactionControllers"
)

func TransactionRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET("/transaction", middleware.Authenticate(), transactionControllers.GetTransactions())
	incomingRoutes.POST("/transaction", middleware.Authenticate(), transactionControllers.AddTransaction())
	incomingRoutes.PATCH("/transaction/:transaction_id", transactionControllers.EditTransaction())
	incomingRoutes.DELETE("/transaction/:transaction_id", transactionControllers.DeleteTransaction())
}
