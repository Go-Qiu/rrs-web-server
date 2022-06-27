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
	"sync"

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

// PointsToVouchers will submit a POST request to:
// (1) vouchers microservice, to issue the request number of vouchers;
// (2) users micrososervice, to deduct the points used from the user's points wallet.
func (u *UserCtl) PointsToVouchers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// gather all the parameters passed in via the request.
	// get the user id
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)

	// extract the payload segment of the JWT passed in via the request header.
	payload, err := utils.GetJWTPayload(r)
	if err != nil {
		utils.SendBadRequestMsgToClient(&w, err)
		return
	}
	fmt.Println(payload)

	// inbound body json.

	inboundBody := struct {
		Points         int             `json:"points"`
		AmountInDollar int             `json:"amount_in_dollar"`
		Vouchers       []utils.Voucher `json:"vouchers"`
	}{}

	err = utils.ParseBody(r, &inboundBody)
	if err != nil {
		customErr := errors.New(`[USERS-CTL] fail to parse JSON body`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// outbound request to vouchers microservice.
	epVouchers := fmt.Sprintf(`%s/getvoucher`, os.Getenv("API_URL_VOUCHERS"))
	optVouchers := utils.RequestOptions{
		API: struct {
			Key      string
			Username string
		}{
			os.Getenv("API_KEY_VOUCHERS"),
			os.Getenv("API_USERNAME_VOUCHERS"),
		},
	}

	// outbound request to users microservice
	// epUsers := fmt.Sprintf(`%s/getvoucher`, os.Getenv("API_URL_USERS"))
	// optUsers := utils.RequestOptions{
	// 	API: struct{Key string; Username string}{
	// 		os.Getenv("API_KEY_USERS"),
	// 		os.Getenv("API_USERNAME_USERS"),
	// 	},
	// }

	// loop through all the requested vouchers (in inbound request), concurrently.
	// applying the multi-producer --> one consumer pattern, via unbuffered channel.

	var wg sync.WaitGroup

	// channel used by producers to send the response out.
	jobs := make(chan utils.RequestOutcome)

	// channel used by consumer to signal it is done and send back the processed data.
	// done := make(chan bool)
	done := make(chan []interface{})

	// loop through the inbound request body to break down all vouchers to qty of 1 pcs.

	// http client to connect to users microservice.
	// setup the client to bypass the ssl verification check so that a call to users microservice (via https, protected by self-signed ssl cert) can be done.
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	vouchers := utils.BreakdwonVouchersToQtyOfOneUnit(inboundBody.Vouchers)

	for _, v := range vouchers {

		pts := strconv.Itoa(v.Points)

		// prepare outbound reqeuest body.
		rb := struct {
			UserID string
			Points string
			Value  string
		}{
			UserID: id,
			Points: pts,
			Value:  v.ValueInDollar,
		}

		reqBody, _ := json.Marshal(rb)

		// prepare the outbound post request.
		apiReq, err := utils.PreparePostRequest(epVouchers, reqBody, optVouchers)
		if err != nil {

			customErr := errors.New("[USERS-CTL] fail to parse request")
			utils.SendErrorMsgToClient(&w, customErr)

			// exit for loop
			break
		}

		// concurrent sending of Post request to vouchers microservice.
		// increment the go routine wait for completion count.
		wg.Add(1)

		// send to producer for execution concurrently.
		go utils.PostRequest(client, apiReq, jobs, &wg)
	}

	// fire up the response consumer.
	go utils.ResponseConsumer(&w, jobs, done)

	// wait for all the POST request producers to complete.
	wg.Wait()

	// POST request producers are done. close the jobs channel.
	close(jobs)

	// block until response consumer is done and received a true from the done channel.
	<-done

	// generate the vouchers in voucher microservice.
	// endpoint_v := fmt.Sprintf(`%s/getvoucher`, API_ROOT_URL)

	// cache the voucher ids returned (for rollback purpose, in exceptions).

	// deduct the points in user microservice.
	// endpoint_u := fmt.Sprintf(`%s/addvoucher`, API_ROOT_URL)

	// conclude the conversion.

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
