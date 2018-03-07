package jsonserver

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var routes = map[string][]Route{}
var routesLock = sync.RWMutex{}

// RegisterRoute stores a closure to execute against a method and path
func RegisterRoute(method string, path string, middleware []Middleware, action RouteAction) {

	methods := strings.Split(method, "|")

	for _, method := range methods {
		routesLock.Lock()
		routes[method] = append(routes[method], Route{Path: path, Action: action, Middleware: middleware})
		routesLock.Unlock()
	}

}

// Dispatch will search for and execute a route
func dispatch(request *http.Request, response *http.ResponseWriter, method string, path string, params string, body *[]byte) (bool, error) {

	routesLock.RLock()

	if methodRoutes, ok := routes[method]; ok {

		routesLock.RUnlock()

		for _, route := range methodRoutes {

			routeMatches, routeParams := route.MatchesPath(path)

			// TODO: Implement a check here that works with (and extracts) wildcards
			if routeMatches {

				queryParams, _ := url.ParseQuery(params)

				for _, middleware := range route.Middleware {

					// Execute all middleware and halt execution if one of them
					// returns FALSE
					if middleware(request, body, queryParams, routeParams) == false {
						return false, errors.New("Access denied to route")
					}

				}

				route.Action(request, response, body, queryParams, routeParams)

				return true, nil

			}

		}

	} else {
		routesLock.RUnlock()
	}

	return false, nil

}
