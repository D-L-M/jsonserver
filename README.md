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

    jsonserver.RegisterRoute("GET", "/url-goes-here", func(request *http.Request, response *http.ResponseWriter, body *[]byte, params url.Values) {

        jsonserver.WriteResponse(response, jsonserver.JSON{"foo":"bar","params":params}, http.StatusOK)

    })

    jsonserver.Start(9999)

    select{}

}
```