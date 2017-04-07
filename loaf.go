package main

import (
	"log"
	"net/http"

	"gitlab.com/infinitewtf/loaf/router"
)

func main() {
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
