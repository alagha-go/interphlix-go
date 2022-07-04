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
	return &oauth2.Token{AccessToken: account.Token.AccessToken, RefreshToken: account.Token.RefreshToken, TokenType: account.Token.TokenType, Expiry: account.Token.Expiry}
}