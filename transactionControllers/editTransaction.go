package transactionControllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swapnika/jwt-auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EditTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		transactionId := c.Param("transaction_id")

		var transaction models.Transaction

		if err := c.BindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D

		transaction.ID, _ = primitive.ObjectIDFromHex(transactionId)
		transaction.Transaction_id = transactionId

		updateObj = append(updateObj, bson.E{Key: "transaction_title", Value: transaction.Transaction_title})
		updateObj = append(updateObj, bson.E{Key: "transaction_amount", Value: transaction.Transaction_amount})

		// transaction.Date = time.Now()
		updateObj = append(updateObj, bson.E{Key: "date", Value: transaction.Date})

		upsert := false
		filter := bson.M{"transaction_id": transactionId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		_, err := transactionCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)
		if err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transaction)
	}

}
