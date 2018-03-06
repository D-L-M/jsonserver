# jsonserver

jsonserver is a simple Golang TCP server and routing component that can be used to create a simple JSON API.

Simple usage is:

```go
package main

import (
    "github.com/D-L-M/jsonserver"
    "net/http"
    "net/url"
)

func main() {

    middleware := []jsonserver.Middleware{} // Optional slice of Middleware functions

    jsonserver.RegisterRoute("GET", "/url-goes-here", middleware, func(request *http.Request, response *http.ResponseWriter, body *[]byte, queryParams url.Values) {

        jsonserver.WriteResponse(response, jsonserver.JSON{"foo": "bar", "query_params": queryParams}, http.StatusOK)

    })

    jsonserver.Start(9999)

    select{}

}
```

Middleware functions have the signature `func(request *http.Request, body *[]byte, queryParams url.Values) bool` and will prevent the route from loading if they return `false`.