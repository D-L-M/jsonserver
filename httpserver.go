package jsonserver

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// HTTP server
type server struct{}

// Handle incoming requests and route to the appropriate package
func (requestHandler *server) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		WriteResponse(response, &JSON{"success": false, "message": "Could not read request body"}, http.StatusBadRequest)
	} else {

		// Write the body back to the request for later use
		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Extract request details and dispatch to the appropriate route
		method := request.Method
		path := request.URL.Path[:]
		params := request.URL.RawQuery
		success, middlewareResponseCode, err := dispatch(request, response, method, path, params, &body)

		// Access denied by middleware
		if err != nil {

			WriteResponse(response, &JSON{"success": false, "message": "Access denied"}, middlewareResponseCode)

			// No matching routes found
		} else if success == false {

			WriteResponse(response, &JSON{"success": false, "message": "Could not find " + path}, http.StatusNotFound)

		}

	}

}

// Start initialises the HTTP server
func Start(port int, timeout int) {

	requestHandler := &server{}
	timeoutDuration := time.Duration(time.Duration(timeout) * time.Second)

	http.Handle("/", http.TimeoutHandler(requestHandler, timeoutDuration, "Request timed out"))

	go func() {

		err := http.ListenAndServe(":"+strconv.Itoa(port), nil)

		if err != nil {
			log.Fatal(err)
		}

	}()

}
