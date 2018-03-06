package jsonserver

import (
	"net/http"
	"net/url"
	"strings"
)

// Route structs define executable HTTP routes
type Route struct {
	Path  string
	Route func(request *http.Request, response *http.ResponseWriter, body *[]byte, routeParams url.Values)
}

var routes = map[string][]Route{}

// RegisterRoute stores a closure to execute against a method and path
func RegisterRoute(method string, path string, route func(request *http.Request, response *http.ResponseWriter, body *[]byte, routeParams url.Values)) {

	methods := strings.Split(method, "|")

	for _, method := range methods {
		routes[method] = append(routes[method], Route{Path: path, Route: route})
	}

}

// Dispatch will search for and execute a route
func dispatch(request *http.Request, response *http.ResponseWriter, method string, path string, params string, body *[]byte) bool {

	if methodRoutes, ok := routes[method]; ok {

		for _, route := range methodRoutes {

			routeParams, _ := url.ParseQuery(params)

			// TODO: Implement a check here that works with (and extracts) wildcards
			if route.Path == path || route.Path == "/*" {

				route.Route(request, response, body, routeParams)

				return true

			}

		}

	}

	return false

}
