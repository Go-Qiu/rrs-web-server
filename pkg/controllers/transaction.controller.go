package controllers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-qiu/rrs-web-server/pkg/utils"
)

// TransactCtl is a struct that represents a user controller.
type TransactionCtl struct {
	name      string
	apikey    string
	dataStore map[string]DataPointExtended
}

func NewTransactionCtl(name string, apikey string, ds map[string]DataPointExtended) *TransactionCtl {
	return &TransactionCtl{
		name:      name,
		apikey:    apikey,
		dataStore: ds,
	}
}

// Create executes the flow to add a recyclable transaction and assocate it to a specific user.
func (t *TransactionCtl) Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// extract the payload segement of the JWT passed in via the request header.
	payload, err := utils.GetJWTPayload(r)
	if err != nil {
		utils.SendBadRequestMsgToClient(&w, err)
		return
	}

	inboundBody := struct {
		ItemCat string  `json:"item_cat"`
		Points  float64 `json:"points"`
		Weight  float64 `json:"wgt_in_grams"`
	}{}

	err = utils.ParseBody(r, &inboundBody)
	if err != nil {
		customErr := errors.New(`[TRANX-CTL] fail to parse JSON body`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// prepare the outbound post request to the users microservice.

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

	// set the endpoint query string
	endpoint := fmt.Sprintf(`%s/addtransaction`, API_ROOT_URL)

	// POST request body json to be send to microservice.
	outboundInputs := Transaction{
		Item:   inboundBody.ItemCat,
		Phone:  payload.Phone,
		Points: inboundBody.Points,
		Weight: inboundBody.Weight,
	}

	reqBody, err := json.Marshal(outboundInputs)
	if err != nil {
		customErr := errors.New(`[TRANX-CTL] fail to parse request body`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// prepare the outbound POST request to the users microservice.
	apiReq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		customErr := errors.New(`[TRANX-CTL] transaction data parsing, failed`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}
	// set all the required header attributes of this POST request.
	apiReq.Header.Set("Content-Type", "application/json")
	apiReq.Header.Set("apiKey", API_KEY)
	apiReq.Header.Set("username", API_USERNAME)

	// send out the outbound post request.
	outcome, err := client.Do(apiReq)
	if err != nil {
		customErr := errors.New(`[TRANX-CTL] transaction submission, failed`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// handle the response from the users microservice.
	var outcomeRespBody ResponseBody
	err = utils.ParseResponseBody(outcome, &outcomeRespBody)
	if err != nil {
		customErr := errors.New(`[TRANX-CTL] transaction submission, failed`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	if !outcomeRespBody.Ok {
		customErr := errors.New(outcomeRespBody.Msg)
		utils.SendBadRequestMsgToClient(&w, customErr)
		return
	}

	// ok.  parse the map to json string.
	data, err := json.Marshal(outcomeRespBody.Data)
	if err != nil {
		customErr := errors.New(`[TRANX-CTL] fail to parse the response from the microservice`)
		utils.SendBadRequestMsgToClient(&w, customErr)
		return
	}

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-USERS]: transaction added and associated to user, successful",
		"data" : %s
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
}
