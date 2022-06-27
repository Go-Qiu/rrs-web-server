package controllers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-qiu/rrs-web-server/pkg/utils"
	"github.com/gorilla/mux"
)

// UserCtl is a struct that represents a user controller.
type UserCtl struct {
	name      string
	apikey    string
	dataStore map[string]DataPointExtended
	CRUD
}

func NewUserCtl(name string, apikey string, ds map[string]DataPointExtended) *UserCtl {
	return &UserCtl{
		name:      name,
		apikey:    apikey,
		dataStore: ds,
	}
}

// AddTransaction executes the flow to add a recyclable transaction and assocate it to a specific user.
func (u *UserCtl) Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-USERS]: added a transaction, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
}

// GetTransactionsByType retrieves all the recyclable transactions, of a specific type (code) that are associated to the user.
func (u *UserCtl) GetTransactionsByType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get the user id
	params := mux.Vars(r)
	id := params["id"]
	typeCode := params["type_code"]

	// prepare the outbound GET request to the users microservice.
	// http client to connect to users microservice.
	// setup the client to bypass the ssl verification check so that a call to users microservice (via https, protected by self-signed ssl cert) can be done.
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// get environment variables for connecting to user microservice.
	API_ROOT_URL := os.Getenv("API_URL_USERS")
	API_KEY := os.Getenv("API_KEY_USERS")
	API_USERNAME := os.Getenv("API_USERNAME_USERS")

	DATA_PTS_PER_PAGE, err := strconv.Atoi(os.Getenv("DATA_PTS_PER_PAGE"))
	if err != nil {
		log.Fatal(err)
	}

	// loop to request for all transactions data, of a type (by page index)
	pageIndex := 0
	dataPtsCount := 0
	isOk := false

	// transactions := []TransactionExpended{}
	transactions := []interface{}{}

	for (pageIndex == 0) || (isOk && dataPtsCount == DATA_PTS_PER_PAGE) {

		// set the endpoint query string
		endpoint := fmt.Sprintf(`%s/gettransactions/%s/%s?page_index=%d&records_per_page=%d`, API_ROOT_URL, id, typeCode, pageIndex, DATA_PTS_PER_PAGE)

		// prepare the outbound GET request to the users microservice.
		apiReq, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			customErr := errors.New(`[USERS-CTL] fail to parse request`)
			utils.SendErrorMsgToClient(&w, customErr)
			return
		}

		// set all the required header attributes of this GET request.
		apiReq.Header.Set("Content-Type", "application/json")
		apiReq.Header.Set("apiKey", API_KEY)
		apiReq.Header.Set("username", API_USERNAME)

		// send out the request.
		outcome, err := client.Do(apiReq)
		if err != nil {
			customErr := errors.New(`[USERS-CTL] fail to submit request`)
			utils.SendErrorMsgToClient(&w, customErr)
			return
		}

		// handle the response body returned by the users microservice.
		type respBody struct {
			Ok   bool                     `json:"ok"`
			Msg  string                   `json:"msg"`
			Data map[string][]interface{} `json:"data"`
		}
		var outcomeRespBody respBody
		err = utils.ParseResponseBody(outcome, &outcomeRespBody)
		if err != nil {
			customErr := errors.New(`[USERS-CTL] get transaction request sumbission, failed`)
			utils.SendErrorMsgToClient(&w, customErr)
			return
		}

		// if !outcomeRespBody.Ok {
		// 	customErr := errors.New(outcomeRespBody.Msg)
		// 	utils.SendBadRequestMsgToClient(&w, customErr)
		// 	return
		// }
		isOk = outcomeRespBody.Ok
		dataPtsCount = 0

		if !outcomeRespBody.Ok && (outcome.StatusCode == http.StatusNotFound) {
			// no data points found.
			customErr := errors.New(outcomeRespBody.Msg)
			utils.SendNotFoundMsgToClient(&w, customErr)

			// exit the for loop
			break
		}

		for _, data := range outcomeRespBody.Data {

			// breakdown the first level map.
			// dp := TransactionExpended{}

			// loop through the second level map.
			// get the attribute and its value.
			// build the user data point struct.
			for _, v := range data {
				// jsonString, _ := json.Marshal(v)
				transactions = append(transactions, v)
				// transactions = append(transactions, dp)
				// increment the number of data points count.
				dataPtsCount++
				//
			}
			//
		}

		// increment the page index
		pageIndex++
		//
	}

	if isOk {
		// ok.
		// data in json string format
		data, err := json.Marshal(transactions)
		if err != nil {
			customErr := errors.New(`[USERS-CTL] parsing returned transactions, failed`)
			utils.SendErrorMsgToClient(&w, customErr)
			return
		}

		utils.SendDataToClient(&w, data, `[USERS-CTL] retrieval of recyle transactions associated to user, successful`)
	}

}

func (u *UserCtl) GetPoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// extract the payload segement of the JWT passed in via the request header.
	payload, err := utils.GetJWTPayload(r)
	if err != nil {
		utils.SendBadRequestMsgToClient(&w, err)
		return
	}

	// prepare the outbound GET request to the users microservice.
	// http client to connect to users microservice.
	// setup the client to bypass the ssl verification check so that a call to users microservice (via https, protected by self-signed ssl cert) can be done.
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// get environment variables for connecting to user microservice.
	API_ROOT_URL := os.Getenv("API_URL_USERS")
	API_KEY := os.Getenv("API_KEY_USERS")
	API_USERNAME := os.Getenv("API_USERNAME_USERS")

	// POST request body json to be send to microservice.
	// reqBody struct to facilitate INBOUND request body json unmarshalling.
	type postInput struct {
		Phone string `json:"phone"`
	}

	sendInput := postInput{
		Phone: payload.Phone,
	}

	reqBody, err := json.Marshal(sendInput)

	// set the endpoint query string
	endpoint := fmt.Sprintf(`%s/retrieve`, API_ROOT_URL)

	// prepare the outbound GET request to the users microservice.
	apiReq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		customErr := errors.New(`[USERS-CTL] fail to parse request`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// set all the required header attributes of this GET request.
	apiReq.Header.Set("Content-Type", "application/json")
	apiReq.Header.Set("apiKey", API_KEY)
	apiReq.Header.Set("username", API_USERNAME)

	// send out the request.
	outcome, err := client.Do(apiReq)
	if err != nil {
		customErr := errors.New(`[USERS-CTL] fail to submit request`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// handle the response body returned by the users microservice.
	type respBody struct {
		Ok   bool                   `json:"ok"`
		Msg  string                 `json:"msg"`
		Data map[string]interface{} `json:"data"`
	}
	var outcomeRespBody respBody
	err = utils.ParseResponseBody(outcome, &outcomeRespBody)
	if err != nil {
		customErr := errors.New(`[USERS-CTL] get transaction request sumbission, failed`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// dataPoint struct to holds the data returned by users microservice.
	dp := struct {
		Points float64
		ID     float64
	}{}

	// break down the map
	for k, v := range outcomeRespBody.Data {

		switch k {
		case "userID":
			dp.ID = v.(float64)
		case "points":
			dp.Points = v.(float64)
		}

	}

	// data in json string format
	data, _ := json.Marshal(dp)

	utils.SendDataToClient(&w, data, `[USERS-CTL] retrieved points collected by user, successful`)

}

func (u *UserCtl) GetVouchers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-USERS]: retrieval of points collected by user, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
}
