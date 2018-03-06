package jsonserver

import (
	"net/http"
	"net/url"
)

// JSON represents JSON documents in map form
type JSON map[string]interface{}

// RouteAction is a function signature for actions carried out when a route is matched
type RouteAction func(request *http.Request, response *http.ResponseWriter, body *[]byte, queryParams url.Values)

// Route structs define executable HTTP routes
type Route struct {
	Path       string
	Action     RouteAction
	Middleware []Middleware
}

// Middleware is a function signature for HTTP middleware that can be assigned routes
type Middleware func(request *http.Request, body *[]byte, queryParams url.Values) bool
