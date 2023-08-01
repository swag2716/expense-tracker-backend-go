package transactionControllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swapnika/jwt-auth/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		userId, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User id not found"})
			return

		}

		var transaction models.Transaction

		if err := c.BindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(transaction)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		transaction.ID = primitive.NewObjectID()
		transaction.Transaction_id = transaction.ID.Hex()

		transaction.User_id = userId.(string)

		_, insertErr := transactionCollection.InsertOne(ctx, transaction)

		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}

		c.JSON(http.StatusOK, transaction)

	}
}
