package accounts

import (
	"errors"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)


func ValidateRequest(req *http.Request) error {
	cookie, err := req.Cookie("token")
	if err != nil {
		return errors.New(variables.NoToken)
	}
	valid, _ := VerifyToken(cookie.Value)
	if !valid {
		return errors.New(variables.InvalidToken)
	}
	return nil
}

func GetAccount(tokenString string) (accounts.Account, error) {
	claims := &Claims{}

	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		secret := variables.LoadSecret()
		return []byte(secret.JwtKey), nil
	})
	account, err := accounts.GetAccountByID(claims.AccountID)
	return account, err
}


func GetmyAccount(req *http.Request) (accounts.Account, error) {
	cookie, err := req.Cookie("token")
	if err != nil {
		return accounts.Account{}, errors.New(variables.NoToken)
	}
	return GetAccount(cookie.Value)
}