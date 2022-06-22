package accounts

import (
	"errors"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)


func ValidateRequest(req *http.Request) bool {
	cookie, err := req.Cookie("token")
	if err != nil {
		return false
	}
	valid, _ := VerifyToken(cookie.Value)
	return valid
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


func GetMyAccount(req *http.Request) (accounts.Account, error) {
	cookie, err := req.Cookie("token")
	if err != nil {
		return accounts.Account{}, errors.New("provide authorization token")
	}
	return GetAccount(cookie.Value)
}