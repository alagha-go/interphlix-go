package handler

import (
	"interphlix/lib/handler/accounts"
	"net/http"

	"github.com/gorilla/mux"
)


var (
	Router = mux.NewRouter()
)

func Main() {
	// routes to work on account
	Router.HandleFunc("/apis/sign-up", accounts.SignUp).Methods("POST")
	Router.HandleFunc("/apis/myaccount", accounts.GetMyAccount).Methods("GET")
}


func ServeFavicon(res http.ResponseWriter, req *http.Request) {
	path := "./web/logo.ico"
	http.ServeFile(res, req, path)
}