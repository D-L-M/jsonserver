package jsonserver

import (
	"io/ioutil"
	"net/http"
	"testing"
)

// TestEnableTLSFailsWithBadCertPath tests that the server fails to start if an
// invalid cert path is provided
func TestEnableTLSFailsWithBadCertPath(t *testing.T) {

	if NewServer().EnableTLS("foo", "./test.key") == nil {
		t.Errorf("Server unexpectedly started")
	}

}

// TestEnableTLSFailsWithBadKeyPath tests that the server fails to start if an
// invalid key path is provided
func TestEnableTLSFailsWithBadKeyPath(t *testing.T) {

	if NewServer().EnableTLS("./test.crt", "foo") == nil {
		t.Errorf("Server unexpectedly started")
	}

}

// TestServerCanListenHTTPS tests starting and making a HTTPS request to the
// server
func TestServerCanListenHTTPS(t *testing.T) {

	testRouteSetUp()

	response, err := http.Get("https://127.0.0.1:9999/")

	if err != nil {

		t.Errorf("Unable to make request")

	} else {

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)

		if err != nil {

			t.Errorf("Unexpected error thrown when attempting to read response")

		} else {

			bodyString := string(body)

			if bodyString != "GET /" {
				t.Errorf("Could not reach route")
			}

		}

	}

	testRouteTearDown()

}

// TestServerCanListenHTTP tests starting and making a HTTP request to the
// server
func TestServerCanListenHTTP(t *testing.T) {

	testRouteSetUp()

	response, err := http.Get("http://127.0.0.1:9998/")

	if err != nil {

		t.Errorf("Unable to make request")

	} else {

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)

		if err != nil {

			t.Errorf("Unexpected error thrown when attempting to read response")

		} else {

			bodyString := string(body)

			if bodyString != "GET /" {
				t.Errorf("Could not reach route")
			}

		}

	}

	testRouteTearDown()

}

// TestServerTimesOutHTTPS tests starting and making a HTTPS request to the
// server that times out
func TestServerTimesOutHTTPS(t *testing.T) {

	testRouteSetUp()

	response, _ := http.Get("https://127.0.0.1:9999/timeout")

	if response.StatusCode != 503 {
		t.Errorf("Request did not timeout like expected")
	}

	testRouteTearDown()

}

// TestServerTimesOutHTTP tests starting and making a HTTP request to the
// server that times out
func TestServerTimesOutHTTP(t *testing.T) {

	testRouteSetUp()

	response, _ := http.Get("http://127.0.0.1:9998/timeout")

	if response.StatusCode != 503 {
		t.Errorf("Request did not timeout like expected")
	}

	testRouteTearDown()

}

// TestServerReturnsNotFound tests receiving a 404 response from the server for a bad route
func TestServerReturnsNotFound(t *testing.T) {

	testRouteSetUp()

	response, err := http.Get("https://127.0.0.1:9999/404")

	if err != nil {

		t.Errorf("Unable to make request")

	} else {

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		code := response.StatusCode

		if err != nil {

			t.Errorf("Unexpected error thrown when attempting to read response")

		} else {

			bodyString := string(body)

			if bodyString != `{"message":"Could not find /404","success":false}` {
				t.Errorf("Route did not return 'not found' message")
			}

			if code != 404 {
				t.Errorf("Route did not return 404 HTTP code")
			}

		}

	}

	testRouteTearDown()

}

// TestServerReturnsOutputFromDenyingMiddleware tests receiving a middleware denial response
func TestServerReturnsOutputFromDenyingMiddleware(t *testing.T) {

	testRouteSetUp()

	response, err := http.Get("https://127.0.0.1:9999/middleware_deny")

	if err != nil {

		t.Errorf("Unable to make request")

	} else {

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		code := response.StatusCode

		if err != nil {

			t.Errorf("Unexpected error thrown when attempting to read response")

		} else {

			bodyString := string(body)

			if bodyString != `{"message":"Access denied","success":false}` {
				t.Errorf("Route did not return middleware denial message")
			}

			if code != 401 {
				t.Errorf("Route did not return middleware denial HTTP code")
			}

		}

	}

	testRouteTearDown()

}
