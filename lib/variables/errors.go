package variables

import "context"


const (
	UserAlreadyExists	= "user with this email already exists"
	UserNotFound = "user not found"
	InternalServerError = "internal server error"
	InvalidJson = "invalid json data"
	InvalidToken = "invalid token"
	NoToken = "no token provided"
	CouldNotGetToken = "error could not get token from google"
	NoCode = "No Google code provided"
	CouldNotGetUserInfoFromGoogle = "could not get your profile info from google"
	CouldNotGenerateToken = "error could not genreare auth token try to relogin later"
	WrongPassword = "wrong password"
	ShortPassword = "your password is too short"
	ProjectExists = "project already exists"
	EmailNotVerified = "email not verified, verify your email and try again"
	InvalidName = "empty name not allowed"
	InvalidID = "invalid id"
	ProjectNotFound = "project does not exist"
	ApiKeyExists = "api key with this name already exists"
	ProjectsLimit = "you cannot create more project delete one to create another one"
	ApiKeysLimit = "reached api key limit for this project delete one to create another one"
	InvalidEmail = "provide a valid email address to proceed"
	MovieExists = "movie with this code already exists"
	MovieNotFound = "movie does not exists"
	SeasonNotFound = "season with does not exist"
)

func SaveError(err error, pkg string, function string) {
	if err != nil {
		ctx := context.Background()
		collection := Local.Database("Interphlix").Collection("Errors")
		log := Log{Error: err.Error(), Package: pkg, Function: function}
		_, err := collection.InsertOne(ctx, log)
		HandleError(err)
	}
}