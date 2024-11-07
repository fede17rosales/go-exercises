package database

import (
	"context"
	"twittergo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckUser(email string) (models.User, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID

}
