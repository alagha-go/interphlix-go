package accounts

import (
	"errors"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	AccountID					primitive.ObjectID
	jwt.StandardClaims
}

func GenerateToken(account accounts.Account) (*http.Cookie, int, error) {
	expires := time.Now().Add(120*time.Hour)
	claims := &Claims{
		AccountID: account.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := variables.LoadSecret()
	tokenString, err := token.SignedString([]byte(secret.JwtKey))
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not generate token")
	}
	return &http.Cookie{Name: "token", Value: tokenString, Domain: ".interphlix.com", Path: "/"}, http.StatusOK, nil
}


func VerifyToken(tokenString string) (bool, int) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		secret := variables.LoadSecret()
		return []byte(secret.JwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, http.StatusUnauthorized
		}
		return false, http.StatusBadRequest
	}

	if !token.Valid {
		return false, http.StatusUnauthorized
	}

	return true, http.StatusOK
}


func RefreshToken(tokenString string) (*http.Cookie, int, error) {
	valid, status := VerifyToken(tokenString)
	if !valid {
		return nil, status, errors.New(variables.InvalidToken)
	}
	account, err := GetAccount(tokenString)
	if err != nil {
		return nil, http.StatusNotFound, errors.New(variables.UserNotFound)
	}
	return GenerateToken(account)
}