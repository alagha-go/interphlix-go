package accounts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)


func ChangePassword(email, password, newPassword string) ([]byte, int) {
	var account Account
	var Response variables.Response
	Response.Action = variables.ChangePassword
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&account)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.UserNotFound
		return variables.JsonMarshal(Response), http.StatusNotFound
	}
	valid := CompareHash(password, account.Password)
	if !valid {
		Response.Failed = true
		Response.Error = variables.WrongPassword
		return variables.JsonMarshal(Response), http.StatusUnauthorized
	}
	hash := Hasher([]byte(newPassword))
	filter := bson.M{"email": bson.M{"$eq": email}}
	update := bson.M{"$set": bson.M{"password": hash}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	Response.Success = true
	return variables.JsonMarshal(Response), http.StatusOK
}