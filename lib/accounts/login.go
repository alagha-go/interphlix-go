package accounts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)


func (account *Account) Login() ([]byte, int) {
	var Account Account
	response := variables.Response{Action:variables.Login}
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")

	err := collection.FindOne(ctx, bson.M{"email": account.Email}).Decode(&Account)
	if err != nil {
		response.Failed = true
		response.Error = variables.UserNotFound
		return variables.JsonMarshal(response), http.StatusNotFound
	}
	account.ID = Account.ID
	valid := CompareHash(account.Password, Account.Password)
	if !valid {
		response.Failed = true
		response.Error = variables.WrongPassword
		return variables.JsonMarshal(response), http.StatusUnauthorized
	}
	response.Success = true
	return variables.JsonMarshal(response), http.StatusOK
}