package utils

import (
	"bytes"
	"crypto/tls"
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
	Outcome *ResponseBody
}

type PointsToVouchersTransaction struct {
	VID           string  `json:"vid"`
	UserID        string  `json:"userid"`
	Points        float64 `json:"points"`
	ValueInDollar string  `json:"value_in_dollar"`
	CreatedDate   string  `json:"created_date"`
}

type ResponseBody struct {
	Ok   bool                   `json:"ok"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type Frame struct {
	Ok   bool
	Msg  string
	Data []PointsToVouchersTransaction
}

// PostRequest will send a post request, concurrently using the passed in wait group. The response body has a data type of map[string]interface{}.
func PostRequest(r *http.Request, jobs chan<- RequestOutcome, wg *sync.WaitGroup) {

	defer wg.Done()
	resp := RequestOutcome{}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// send out a POST request to the microservices.
	outcome, err := client.Do(r)
	if err != nil || (outcome.StatusCode != http.StatusOK) {
		// post request failed.
		resp.Ok = false
		resp.Msg = "[CON] send post request, failed"
		resp.Outcome = nil

		// send the response into the channel.
		jobs <- resp
		return
	}

	// ok.
	outcomeBody := ResponseBody{}
	err = ParseResponseBody(outcome, &outcomeBody)
	if err != nil {
		resp.Ok = false
		resp.Msg = `[CON] fail to parse response`
		resp.Outcome = nil
	}

	resp.Ok = true
	resp.Msg = "[CONC] send post request, successful"
	resp.Outcome = &outcomeBody

	// send the response into the channel.
	jobs <- resp
	//
}

// FetchRequest will send a GET request, concurrently using the passed in wait group.
func FetchRequest(r *http.Request, wg *sync.WaitGroup) {

	//
}

// ResponseConsumer will handle the response returned via the jobs channel (unbuffered).
func ResponseConsumer(w *http.ResponseWriter, jobs <-chan RequestOutcome, done chan<- Frame) {

	// flag to indicate at least 1 error has occurred.
	hasError := false

	// cache to harvest all the response data received.
	dataPoints := []PointsToVouchersTransaction{}

	// handle the request outcome received via the jobs channel.
	for ro := range jobs {

		// failure in posting request.
		if !ro.Ok {
			// error handling here.
			hasError = true
			continue
		}

		if !ro.Outcome.Ok {
			hasError = true
			continue
		}

		// ok. append the data into the existing dataPoints slice.
		// dataPoints = append(dataPoints, ro.Outcome.Data...)

		// breakdown the map[string]interface{}
		dp := PointsToVouchersTransaction{}

		for k, v := range ro.Outcome.Data {
			switch k {
			case "VID":
				dp.VID = v.(string)
			case "UserID":
				dp.UserID = v.(string)
			case "Pionts":
				dp.Points = v.(float64)
			case "Value":
				dp.ValueInDollar = v.(string)
			case "CreatedDate":
				dp.CreatedDate = v.(string)
			}

		}
		dataPoints = append(dataPoints, dp)

	}

	if hasError {
		// error handling
		rb := Frame{
			Ok:   false,
			Msg:  `[CONC] fail to parse response`,
			Data: dataPoints,
		}
		done <- rb
		return
	}

	// ok.
	rb := Frame{
		Ok:   true,
		Msg:  `[CONC] success`,
		Data: dataPoints,
	}

	done <- rb
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
	apiReq.Header.Set("Key", options.API.Key)
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
