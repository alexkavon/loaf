package router

import (
	"gitlab.com/alexkavon/loaf/controllers"
)

func init() {
	routes = append(routes, &Route{name: "home", path: "/", HandlerFunc: controllers.Feed})
	routes = append(routes, &Route{name: "discover", path: "/discover", HandlerFunc: controllers.Discover})
}
