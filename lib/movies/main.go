package movies

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	IndexMovies()
}

func HandleError(err error) {
	if err != nil && err != mongo.ErrEmptySlice{
		log.Panic(err)
	}
}