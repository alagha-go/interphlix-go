package accounts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"interphlix/lib/variables"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	clientfile = "./client.json"
	scopes = []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/drive.file"}
)

/// get account by providing id
func GetAccountByID(ID primitive.ObjectID, full ...bool) (Account, error) {
	var opts *options.FindOneOptions
	var account Account
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")
	if len(full) > 0 && full[0] {
		}else {
			opts = options.FindOne().SetProjection(bson.D{{"password", 0}, {"token", 0}})
	}
	err := collection.FindOne(ctx, bson.M{"_id": ID}, opts).Decode(&account)
	if err != nil {
		return account, errors.New(variables.UserNotFound)
	}
	return account, nil
}


// get account by email 
func GetAccountByEmail(email string, full ...bool) (Account, error) {
	var opts *options.FindOneOptions
	var account Account
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")
	if len(full) > 0 && full[0] {
		}else {
			opts = options.FindOne().SetProjection(bson.D{{"password", 0}, {"token", 0}})
	}
	err := collection.FindOne(ctx, bson.M{"email": email}, opts).Decode(&account)
	if err != nil {
		return account, errors.New(variables.UserNotFound)
	}
	return account, nil
}


// get google configuration
func GetConfig() (*oauth2.Config, error) {
	secretBody, err := ioutil.ReadFile(clientfile)
	if err != nil {
		return nil, err
	}
	return google.ConfigFromJSON(secretBody, scopes...)
}


/// get account info from google
func GetUserInfo(token string) (Account, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", token), nil)
	if err != nil {
		return Account{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return Account{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Account{}, err
	}
	var account GoogleAccount
	err = json.Unmarshal(body, &account)
	if err != nil {
		return Account{}, err
	}
	return Account{Email: account.Email, EmailVerified: account.EmailVerified, UserName: account.Name, FirstName: account.GivenName, LastName: account.FamilyName, Photo: account.Picture, Locale: account.Locale}, nil
}