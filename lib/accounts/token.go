package accounts

import (
	"encoding/json"

	"golang.org/x/oauth2"
)

func (account *Account) SetToken(token *oauth2.Token) {
	data, _ := json.Marshal(token)
	json.Unmarshal(data, &account.Token)
}

func (account *Account) GetToken() *oauth2.Token {
	var token oauth2.Token
	data, _ := json.Marshal(account.Token)
	json.Unmarshal(data, &token)
	return &token
}