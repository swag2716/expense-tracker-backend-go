package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swapnika/jwt-auth/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// extracts the value of user_id parameter in URL path using params
		userId := c.Param("user_id")

		// to perform operations under deadline
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// find the document with the specified user_id
		var user models.SignUpUser
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
