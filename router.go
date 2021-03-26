package jsonserver

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// Router represents an instance of a router
type Router struct {
	Routes     map[string][]Route
	RoutesLock sync.RWMutex
}

// RegisterRoute stores a closure to execute against a method and path
func (router *Router) RegisterRoute(method string, path string, middleware []Middleware, action RouteAction) {

	methods := strings.Split(strings.ToUpper(method), "|")

	for _, method := range methods {

		router.RoutesLock.Lock()

		if router.Routes == nil {
			router.Routes = map[string][]Route{}
		}

		router.Routes[method] = append(router.Routes[method], Route{Path: path, Action: action, Middleware: middleware})

		router.RoutesLock.Unlock()

	}

}

// Dispatch will search for and execute a route
func (router *Router) Dispatch(request *http.Request, response http.ResponseWriter, method string, path string, params string, body *[]byte) (bool, int, error) {

	router.RoutesLock.RLock()

	if methodRoutes, ok := router.Routes[strings.ToUpper(method)]; ok {

		router.RoutesLock.RUnlock()

		for _, route := range methodRoutes {

			routeMatches, routeParams := route.MatchesPath(path)

			if routeMatches {

				queryParams, _ := url.ParseQuery(params)

				ctx := context.Background()
				ctx = context.WithValue(ctx, "state", &RequestState{})
				ctx = context.WithValue(ctx, "routeParams", routeParams)
				ctx = context.WithValue(ctx, "queryParams", &queryParams)

				for _, middleware := range route.Middleware {

					// Execute all middleware and halt execution if one of them
					// returns FALSE
					middlewareDecision, middlewareResponseCode := middleware(ctx, request, response, body)

					if middlewareDecision == false {
						return false, middlewareResponseCode, errors.New("Access denied to route")
					}

				}

				route.Action(ctx, request, response, body)

				return true, 0, nil

			}

		}

	} else {
		router.RoutesLock.RUnlock()
	}

	return false, 0, nil

}
