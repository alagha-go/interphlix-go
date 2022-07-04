package handler

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/handler/projects"

	"github.com/gorilla/mux"
)

var (
	Router = mux.NewRouter()
)

func init() {
	Main()
}

func Main() {
	// routes to work on account
	Router.HandleFunc("/apis/sign-up", accounts.SignUp).Methods("POST")
	Router.HandleFunc("/apis/login", accounts.Login).Methods("POST")
	Router.HandleFunc("/apis/account/changepassword", accounts.ChangePassword).Methods("UPDATE", "PATCH", "PUT")
	Router.HandleFunc("/apis/account/update", accounts.UpdateAccount).Methods("UPDATE", "PATCH", "PUT")
	Router.HandleFunc("/apis/myaccount", accounts.GetMyAccount).Methods("GET")
	Router.HandleFunc("/login/redirect", accounts.Redirect).Methods("GET")
	Router.HandleFunc("/apis/sign-up/google", accounts.GoogleSignUp).Methods("GET")
	Router.HandleFunc("/apis/login/google", accounts.GoogleLogin).Methods("GET")
	
	
	// routes to work on projects
	Router.HandleFunc("/apis/projects/create", projects.CreateProject).Methods("POST")
	Router.HandleFunc("/apis/projects/{projectId}/keys/create", projects.GenerateApiKey).Methods("POST")
	Router.HandleFunc("/apis/projects", projects.GetMyProjects).Methods("GET")
	Router.HandleFunc("/apis/projects/{projectId}/keys", projects.GetProjectApiKeys).Methods("GET")

}