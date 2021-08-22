package jsonserver

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Server represents a HTTP server
type Server struct {
	Router   *Router
	CertPath string
	KeyPath  string
}

// NewServer creates a new server
func NewServer() *Server {

	return &Server{Router: &Router{}}

}

// EnableTLS enabled TLS for the server
func (server *Server) EnableTLS(certPath string, keyPath string) error {

	_, err := os.Stat(certPath)

	if os.IsNotExist(err) {
		return err
	}

	_, err = os.Stat(keyPath)

	if os.IsNotExist(err) {
		return err
	}

	server.CertPath = certPath
	server.KeyPath = keyPath

	return nil

}

// RegisterRoute stores a closure to execute against a method and path
func (server *Server) RegisterRoute(method string, path string, middleware []Middleware, action RouteAction) {

	server.Router.RegisterRoute(method, path, middleware, action)

}

// Handle incoming requests and route to the appropriate package
func (server *Server) ServeHTTP(response http.ResponseWriter, request *http.Request) {

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
		success, middlewareResponseCode, err := server.Router.Dispatch(request, response, method, path, params, &body)

		// Access denied by middleware
		if err != nil {

			WriteResponse(response, &JSON{"success": false, "message": "Access denied"}, middlewareResponseCode)

			// No matching routes found
		} else if !success {

			WriteResponse(response, &JSON{"success": false, "message": "Could not find " + path}, http.StatusNotFound)

		}

	}

}

// Start initialises the HTTP server
func (server *Server) Start(port int, timeout int) {

	timeoutDuration := time.Duration(time.Duration(timeout) * time.Second)
	mux := http.NewServeMux()

	mux.Handle("/", http.TimeoutHandler(server, timeoutDuration, "Request timed out"))

	go func() {

		// HTTPS requests
		if server.CertPath != "" && server.KeyPath != "" {

			certificate, err := tls.LoadX509KeyPair(server.CertPath, server.KeyPath)

			if err != nil {
				log.Fatal(err)
			}

			httpServer := &http.Server{
				Addr:    ":" + strconv.Itoa(port),
				Handler: mux,
				TLSConfig: &tls.Config{
					Certificates: []tls.Certificate{certificate},
					MinVersion:   tls.VersionTLS12,
				},
			}

			err = httpServer.ListenAndServeTLS("", "")

			if err != nil {
				log.Fatal(err)
			}

			// HTTP requests
		} else {

			err := http.ListenAndServe(":"+strconv.Itoa(port), mux)

			if err != nil {
				log.Fatal(err)
			}

		}

	}()

}
