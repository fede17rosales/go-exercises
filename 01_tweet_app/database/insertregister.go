package database

import (
	"context"

	"twittergo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(u models.User) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
