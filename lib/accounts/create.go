package accounts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create account
func CreateAccount(account Account) ([]byte, int) {
	var Response variables.Response
	var accountExist Account
	Response.Action = "create user"
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")

	err := collection.FindOne(ctx, bson.M{"email": account.Email}).Decode(&accountExist)
	if err == nil {
		Response.Failed = true
		Response.Error = variables.UserAlreadyExists
		return variables.JsonMarshal(Response), http.StatusConflict
	}
	account.ID = GetNewAccountID()
	_, err = collection.InsertOne(context.Background(), account)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusConflict
	}
	Response.Data = account
	Response.Success = true
	return variables.JsonMarshal(Response), http.StatusOK
}


/// generate a new account id and make sure it is not duplicate
func GetNewAccountID() primitive.ObjectID {
	ID := primitive.NewObjectID()
	var account Account
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")
	err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&account)
	if err == nil {
		return GetNewAccountID()
	}
	return ID
}