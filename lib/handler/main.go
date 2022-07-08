package handler

import (
	"fmt"
	"interphlix/lib/handler/accounts"
	"interphlix/lib/handler/movies"
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
	fmt.Println("server started successfully")
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
	Router.HandleFunc("/apis/projects/delete/{projectId}", projects.DeleteProject).Methods("DELETE")
	Router.HandleFunc("/apis/projects/{projectId}/keys/delete/{name}", projects.DeleteApiKey).Methods("DELETE")

	// routes to work on movies
	Router.HandleFunc("/apis/movies/upload", movies.UploadMovie).Methods("POST")
	Router.HandleFunc("/apis/movies/{movieId}/seasons/upload", movies.UploadSeason).Methods("POST")
	Router.HandleFunc("/apis/movies/{movieId}/seasons/{seasonId}/episodes/upload", movies.UploadEpisode).Methods("POST")
	Router.HandleFunc("/apis/home", movies.GetRecommendationMovies).Methods("GET")
	Router.HandleFunc("/apis/movies/{id}", movies.GetMovie).Methods("GET")
	Router.HandleFunc("/apis/movies/seasons/{id}/episodes", movies.GetEpisodes).Methods("GET")
	Router.HandleFunc("/apis/genres", movies.GetGenres).Methods("GET")
	Router.HandleFunc("/apis/casts", movies.GetCasts).Methods("GET")
	Router.HandleFunc("/apis/casts/search", movies.SearchCast).Methods("GET")
	Router.HandleFunc("/apis/casts/{id}", movies.GetCast).Methods("GET")
	Router.HandleFunc("/apis/casts/{id}/movies", movies.GetMoviesByCast).Methods("GET")
	Router.HandleFunc("/apis/{type}/search", movies.Search).Methods("GET")
	Router.HandleFunc("/apis/movies/{id}/seasons", movies.GetSeasons).Methods("GET")
	Router.HandleFunc("/apis/{type}/{genre}", movies.GetMovies).Methods("GET")
}
