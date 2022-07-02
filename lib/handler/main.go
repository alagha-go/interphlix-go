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
	fs := http.FileServer(http.Dir("./web/src/"))
	// routes to work on account
	Router.HandleFunc("/apis/sign-up", accounts.SignUp).Methods("POST")
	Router.HandleFunc("/apis/myaccount", accounts.GetMyAccount).Methods("GET")


	// server static files
	Router.HandleFunc("/favicon.ico", ServeFavicon).Methods("GET")
	Router.PathPrefix("/").Handler(fs)
}


func ServeFavicon(res http.ResponseWriter, req *http.Request) {
	path := "./web/logo.ico"
	http.ServeFile(res, req, path)
}