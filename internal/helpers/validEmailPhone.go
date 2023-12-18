package helpers

import (
	"context"

	"github.com/hanshal101/jwt/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = db.GetUsersCollection("users")

var ctx = context.TODO()

func ValidEmailPhone(email string, phone string) (bool, error) {
	filter := bson.M{"email": email, "phone": phone}
	result := collectionName.FindOne(ctx, filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, result.Err()
	}

	return false, nil
}
