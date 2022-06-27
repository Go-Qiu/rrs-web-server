package utils

import (
	"bytes"
	"encoding/json"
	"errors"
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

type RequestOutcome struct {
	Ok      bool
	Msg     string
	Outcome *http.Response
}

type ResponseBody struct {
	Ok   bool        `json:"ok"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// PostRequest will send a post request, concurrently using the passed in wait group.
func PostRequest(c *http.Client, r *http.Request, jobs chan<- RequestOutcome, wg *sync.WaitGroup) {

	resp := RequestOutcome{}

	// send out a POST request to the microservices.
	outcome, err := c.Do(r)
	if err != nil {
		resp.Ok = false
		resp.Msg = err.Error()
		resp.Outcome = nil

		// send the response into the channel.
		jobs <- resp
		return
	}

	// ok.

	resp.Ok = false
	resp.Msg = err.Error()
	resp.Outcome = outcome

	// send the response into the channel.
	jobs <- resp
	//
}

// FetchRequest will send a GET request, concurrently using the passed in wait group.
func FetchRequest(r *http.Request, wg *sync.WaitGroup) {

	//
}

// ResponseConsumer will handle the response returned via the jobs channel (unbuffered).
func ResponseConsumer(w *http.ResponseWriter, jobs <-chan RequestOutcome, done chan<- []interface{}) {

	// flag to indicate at least 1 error has occurred.
	hasError := false

	// cache to harvest all the response data received.
	data := []interface{}{}

	// handle the responses received via the jobs channel.
	for ro := range jobs {

		// failure in posting request.
		if !ro.Ok {
			// error handling here.
			hasError = true
			continue
		}

		// handle the response.
		var outcomeRespBody ResponseBody
		err := ParseResponseBody(ro.Outcome, &outcomeRespBody)
		if err != nil {
			hasError = true
			continue
		}

		if !outcomeRespBody.Ok {
			hasError = true
		}

		// ok.

		dp, err := json.Marshal(outcomeRespBody.Data)
		if err != nil {
			hasError = true
		}
		data = append(data, dp)
	}

	if hasError {
		// error handling
		customErr := errors.New(`[CONC] fail to parse response`)
		SendErrorMsgToClient(w, customErr)
	}

	// send true into the channel, to indicate all responses have been handled with no error.
	done <- data
	//
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
