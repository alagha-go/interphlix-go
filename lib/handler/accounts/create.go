package accounts

import (
	"encoding/json"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"net/http"
)


func SignUp(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var account accounts.Account
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		response := variables.Response{Action: variables.CreateUserAction, Failed: true, Error: variables.InvalidJson}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(response))
		return
	}
	account.EmailVerified = false
	data, status := accounts.CreateAccount(&account)
	if status == 201 {
		token, _, _ := GenerateToken(account)
		http.SetCookie(res, &http.Cookie{
			Domain: ".interphlix.com",
			Name: "token",
			Value: token,
			Path: "/",
			Secure: true,
		})
	}
	res.WriteHeader(status)
	res.Write(data)
}