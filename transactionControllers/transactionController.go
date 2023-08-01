package transactionControllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/swapnika/jwt-auth/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var transactionCollection *mongo.Collection = database.OpenCollection(database.Client, "transaction")
var validate = validator.New()
