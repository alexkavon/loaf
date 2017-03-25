package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	name        string
	path        string
	HandlerFunc http.HandlerFunc
}

var routes []*Route

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, v := range routes {
		router.Path(v.path).Handler(v.HandlerFunc).Name(v.name)
	}

	return router
}
