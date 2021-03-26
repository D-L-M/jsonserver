package jsonserver

import (
	"context"
	"net/http"
)

// JSON represents JSON documents in map form
type JSON map[string]interface{}

// RouteParams is an alias for a map to hold route wildcard parameters, where both keys and values will be strings
type RouteParams map[string]string

// RouteAction is a function signature for actions carried out when a route is matched
type RouteAction func(ctx context.Context, request *http.Request, response http.ResponseWriter, body *[]byte)

// Middleware is a function signature for HTTP middleware that can be assigned to routes
type Middleware func(ctx context.Context, request *http.Request, response http.ResponseWriter, body *[]byte) (bool, int)
