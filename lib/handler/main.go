package handler

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/handler/projects"

	"github.com/gorilla/mux"
)


var (
	Router = mux.NewRouter()
)

func Main() {
	// routes to work on account
	Router.HandleFunc("/apis/sign-up", accounts.SignUp).Methods("POST")
	Router.HandleFunc("/apis/login", accounts.Login).Methods("POST")
	Router.HandleFunc("/apis/projects/create", projects.CreateProject).Methods("POST")
	Router.HandleFunc("/apis/projects/{projectId}/keys/create", projects.GenerateApiKey).Methods("POST")
	Router.HandleFunc("/apis/account/changepassword", accounts.ChangePassword).Methods("POST")
	Router.HandleFunc("/apis/myaccount", accounts.GetMyAccount).Methods("GET")
	Router.HandleFunc("/login/redirect", accounts.Redirect).Methods("GET")
	Router.HandleFunc("/apis/sign-up/google", accounts.GoogleSignUp).Methods("GET")
	Router.HandleFunc("/apis/login/google", accounts.GoogleLogin).Methods("GET")
}