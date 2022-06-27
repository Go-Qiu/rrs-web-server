package utils

import (
	"bytes"
	"net/http"
	"sync"
)

// RequestOptions struct holds additional params needed in an api request.
type RequestOptions struct {
	API struct {
		Key      string
		Username string
	}
}

// PostRequest will send a post request, concurrently using the passed in wait group.
func PostRequest(r *http.Request, wg *sync.WaitGroup) (string, error) {

	return "", nil
}

// FetchRequest will send a GET request, concurrently using the passed in wait group.
func FetchRequest(r *http.Request, wg *sync.WaitGroup) (string, error) {

	return "", nil

}

// PreparePostRequest will prepare a http POST requst for use.
func PreparePostRequest(endpoint string, reqBody []byte, options RequestOptions) (*http.Request, error) {

	apiReq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// set all the required header attributes of this POST request.
	apiReq.Header.Set("Content-Type", "application/json")
	apiReq.Header.Set("apiKey", options.API.Key)
	apiReq.Header.Set("username", options.API.Username)

	return apiReq, nil
}

// PrepareGetRequest will prepare a http GET request for use.
func PrepareGetRequest(endpoint string, options RequestOptions) (*http.Request, error) {

	apiReq, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// set all the required header attributes of this POST request.
	apiReq.Header.Set("Content-Type", "application/json")
	apiReq.Header.Set("apiKey", options.API.Key)
	apiReq.Header.Set("username", options.API.Username)

	return apiReq, nil
}
