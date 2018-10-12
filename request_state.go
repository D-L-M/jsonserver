package jsonserver

// RequestState allows storage of miscellaneous state objects that can be referred to throughout a request
type RequestState struct {
	state map[string]interface{}
}

// Get obtains a state value (or nil if it does not exist)
func (requestState *RequestState) Get(key string) interface{} {

	if value, ok := requestState.state[key]; ok {
		return value
	}

	return nil

}

// Set stores a state value
func (requestState *RequestState) Set(key string, value interface{}) {

	if requestState.state == nil {
		requestState.state = map[string]interface{}{}
	}

	requestState.state[key] = value

}
