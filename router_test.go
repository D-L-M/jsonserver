package jsonserver

import (
	"net/http"
	"net/url"
	"testing"
)

// TestRegisterRoute tests registering a route with the router
func TestRegisterRoute(t *testing.T) {

	RegisterRoute("GET", "/foo", []Middleware{}, func(request *http.Request, response *http.ResponseWriter, body *[]byte, queryParams url.Values, routeParams RouteParams) {
	})

	routesLock.RLock()
	route := routes["GET"][0]
	routesLock.RUnlock()

	if route.Path != "/foo" {
		t.Errorf("Route path mismatch (expected: %v, actual: %v)", "/foo", route.Path)
	}

	if len(route.Middleware) != 0 {
		t.Errorf("Route middleware count mismatch (expected: %v, actual: %v)", 0, len(route.Middleware))
	}

	if route.Action == nil {
		t.Errorf("Route action missing")
	}

	routesLock.Lock()
	routes = map[string][]Route{}
	routesLock.Unlock()

}

// TestRegisterRouteToMultipleMethods tests registering a route with the router against multiple HTTP methods
func TestRegisterRouteToMultipleMethods(t *testing.T) {

	RegisterRoute("GET|PUT", "/bar", []Middleware{func(request *http.Request, body *[]byte, queryParams url.Values, routeParams RouteParams) (bool, int) {
		return false, 401
	}}, func(request *http.Request, response *http.ResponseWriter, body *[]byte, queryParams url.Values, routeParams RouteParams) {
	})

	routesLock.RLock()
	getRoute := routes["GET"][0]
	putRoute := routes["PUT"][0]
	routesLock.RUnlock()

	if getRoute.Path != "/bar" {
		t.Errorf("Route path mismatch (expected: %v, actual: %v)", "/foo", getRoute.Path)
	}

	if len(getRoute.Middleware) != 1 {
		t.Errorf("Route middleware count mismatch (expected: %v, actual: %v)", 1, len(getRoute.Middleware))
	}

	if getRoute.Action == nil {
		t.Errorf("Route action missing")
	}

	if putRoute.Path != "/bar" {
		t.Errorf("Route path mismatch (expected: %v, actual: %v)", "/foo", putRoute.Path)
	}

	if len(putRoute.Middleware) != 1 {
		t.Errorf("Route middleware count mismatch (expected: %v, actual: %v)", 1, len(putRoute.Middleware))
	}

	if putRoute.Action == nil {
		t.Errorf("Route action missing")
	}

	routesLock.Lock()
	routes = map[string][]Route{}
	routesLock.Unlock()

}
