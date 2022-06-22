package accounts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// get account by providing id
func GetAccountByID(ID primitive.ObjectID) (Account, error) {
	var account Account
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")
	err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&account)
	return account, err
}


// get account by email 
func GetAccountByEmail(email string) (Account, error) {
	var account Account
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")
	err := collection.FindOne(ctx, bson.M{"email": account.Email}).Decode(&account)
	return account, err
}