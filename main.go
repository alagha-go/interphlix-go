package main

import (
	_ "interphlix/lib/variables"
	_"interphlix/lib/crawler"
	"interphlix/lib/handler"
	"log"
	"net/http"
)

var (
	PORT = ":8000"
)


func main() {
	err := http.ListenAndServe(PORT, handler.Router)
	HandlError(err)
}


// handle errors by pannic
func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}