package accounts

import (
	"context"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	clientfile = "client.json"
	scopes = []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"}
)


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