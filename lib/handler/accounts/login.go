package accounts

import (
	"encoding/json"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"net/http"
)


func Redirect(res http.ResponseWriter, req *http.Request) {
	var account accounts.Account
	code := req.URL.Query().Get("code")
	if code == "" {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.CreateUserAction, Error: variables.NoCode, Failed: true}))
		return
	}
	token, err := GetToken(code)
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.CreateUserAction, Error: variables.CouldNotGetToken, Failed: true}))
		return
	}
	account, err = GetUserInfo(token.AccessToken)
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.CreateUserAction, Error: variables.CouldNotGetUserInfoFromGoogle, Failed: true}))
		return
	}
	account.SetToken(token)
	data, status := accounts.CreateAccount(&account)
	json.Unmarshal(data, &account)
	cookie, status1, err := GenerateToken(account)
	if err != nil {
		res.WriteHeader(status1)
		res.Write(variables.JsonMarshal(variables.Response{Action: variables.CreateUserAction}))
	}
	http.SetCookie(res, cookie)
	res.WriteHeader(status)
	res.Write(data)
}


func Login(res http.ResponseWriter, req *http.Request) {
	
}