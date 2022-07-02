package handler

import (
	"interphlix/lib/handler/accounts"

	"github.com/gorilla/mux"
)


var (
	Router = mux.NewRouter()
)

func Main() {
	// routes to work on account
	Router.HandleFunc("/apis/sign-up", accounts.SignUp).Methods("POST")
	Router.HandleFunc("/login", accounts.Login).Methods("POST")
	Router.HandleFunc("/apis/myaccount", accounts.GetMyAccount).Methods("GET")
	Router.HandleFunc("/login/redirect", accounts.Redirect).Methods("GET")
}