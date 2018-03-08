# jsonserver

jsonserver is a simple Golang TCP server and routing component that can be used to create a simple JSON API.

Simple usage is:

```go
package main

import (
    "net/http"
    "net/url"

    "github.com/D-L-M/jsonserver"
)

func main() {

    middleware := []jsonserver.Middleware{} // Optional slice of Middleware functions

    jsonserver.RegisterRoute("GET", "/products/{id}", middleware, products)

    jsonserver.Start(9999)

    select{}

}

func products(request *http.Request, response *http.ResponseWriter, body *[]byte, queryParams url.Values, routeParams jsonserver.RouteParams) {

        jsonserver.WriteResponse(*response, jsonserver.JSON{"foo": "bar", "query_params": queryParams, "route_params": routeParams}, http.StatusOK)

    }
```

A route can listen on multiple HTTP methods by pipe-delimiting them, e.g. `GET|POST`.

Named `{wildcard}` fragments in routes are provided in the `routeParams` map, and if a route path ends with `/:` all URL fragments at (and following) that point are collected into a route parameter named `{catchAll}`.

Middleware functions have the signature `func(request *http.Request, body *[]byte, queryParams url.Values, routeParams jsonserver.RouteParams) (bool, int)` and will prevent the route from loading if they return `false` as the boolean value (the int value should be a HTTP status code, which will be used if the middleware returns `false`, or otherwise ignored).