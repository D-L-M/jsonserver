package jsonserver

import (
	"testing"
)

// TestRequestStateGetterAndSetter tests the getter and setter of the RequestState struct
func TestRequestStateGetterAndSetter(t *testing.T) {

	requestState := RequestState{}

	requestState.Set("foo", "bar")

	if requestState.Get("foo") == nil {
		t.Error("Request state not successfully set")
	}

	if requestState.Get("bar") != nil {
		t.Error("Request state unexpectedly set")
	}

}
