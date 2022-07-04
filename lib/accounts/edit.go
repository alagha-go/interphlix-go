package accounts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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


func (account *Account) Update() ([]byte, int) {
	var Response variables.Response
	Response.Action = variables.UpdateAccount
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")

	DBaccount, err := GetAccountByID(account.ID, true)
	if err != nil {
		Response.Failed = true
		Response.Error = err.Error()
		return variables.JsonMarshal(Response), http.StatusNotFound
	}

	if DBaccount.SignUpMethod == "google" {
		return DBaccount.UpdateGoogle()
	}

	filter := bson.M{"_id": bson.M{"$eq": account.ID}}
	m := make(map[string]interface{})
	if account.Email != "" {
		if !EmailValid(account.Email) {
			Response.Failed = true
			Response.Error = variables.InvalidEmail
			return variables.JsonMarshal(Response), http.StatusBadRequest
		}
		_, err := GetAccountByEmail(account.Email)
		if err == nil {
			Response.Failed = true
			Response.Error = variables.UserAlreadyExists
			return variables.JsonMarshal(Response), http.StatusConflict
		}
		m["email"] = account.Email
		m["email_verified"] = false
	}
	
	if account.UserName != "" {
		m["user_name"] = account.UserName
	}
	if account.FirstName != "" {
		m["first_name"] = account.FirstName
	}
	if account.LastName != "" {
		m["last_name"] = account.LastName
	}

	update := bson.M{"$set": m}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNilDocument {
			Response.Failed = true
			Response.Error = variables.UserNotFound
			return variables.JsonMarshal(Response), http.StatusNotFound
		}
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}

	Response.Success = true
	Response.Data = "account updated"
	return variables.JsonMarshal(Response),  http.StatusOK
}

// update account info from google
func (account *Account) UpdateGoogle() ([]byte, int) {
	var Response variables.Response
	Response.Action = variables.UpdateAccount
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")
	token := account.GetToken()
	config, err := GetConfig()
	if err != nil {
		variables.SaveError(err, "accounts", "account.UpdateGoogle")
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	if !token.Valid() {
		src := config.TokenSource(ctx, token)
		token, err = src.Token()
		if err != nil {
			variables.SaveError(err, "accounts", "account.UpdateGoogle")
			Response.Failed = true
			Response.Error = variables.InternalServerError
			return variables.JsonMarshal(Response), http.StatusInternalServerError
		}
	}

	newAccount, err := GetUserInfo(token.AccessToken)
	if err != nil {
		variables.SaveError(err, "accounts", "account.UpdateGoogle")
		Response.Failed = true
		Response.Error = variables.CouldNotGetUserInfoFromGoogle
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}
	newAccount.ID = account.ID
	filter := bson.M{"_id": bson.M{"$eq": newAccount.ID}}
	update := bson.M{"$set": newAccount}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		Response.Failed = true
		Response.Error = variables.InternalServerError
		return variables.JsonMarshal(Response), http.StatusInternalServerError
	}

	Response.Success = true
	Response.Data = "account updated"
	return variables.JsonMarshal(Response),  http.StatusOK
}


///  verify email address
func EmailValid(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}