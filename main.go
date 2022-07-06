package main

import (
	"interphlix/lib/handler"
	_ "interphlix/lib/variables"
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