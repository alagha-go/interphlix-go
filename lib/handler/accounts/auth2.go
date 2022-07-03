package accounts

import (
	"context"
	"encoding/json"
	"fmt"
	"interphlix/lib/accounts"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	clientfile = "client.json"
)

var (
	scopes = []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/drive.file"}
)

func GetUserInfo(token string) (accounts.Account, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", token), nil)
	if err != nil {
		return accounts.Account{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return accounts.Account{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return accounts.Account{}, err
	}
	var account accounts.GoogleAccount
	err = json.Unmarshal(body, &account)
	if err != nil {
		return accounts.Account{}, err
	}
	return accounts.Account{Email: account.Email, EmailVerified: account.EmailVerified, UserName: account.Name, FirstName: account.GivenName, LastName: account.FamilyName, Photo: account.Picture, Locale: account.Locale}, nil
}


func GetToken(code string) (*oauth2.Token, error) {
	config, err := GetConfig()
	if err != nil {
		return nil, err
	}
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}


func GetConfig() (*oauth2.Config, error) {
	secretBody, err := ioutil.ReadFile(clientfile)
	if err != nil {
		return nil, err
	}
	return google.ConfigFromJSON(secretBody, scopes...)
}

/// gets google login url
func GetUrl() (string, error) {
	config, err := GetConfig()
	if err != nil {
		return "", err
	}
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return url, nil
}