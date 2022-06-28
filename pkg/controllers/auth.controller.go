package controllers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-qiu/rrs-web-server/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

// JWTConfig is a struct for storing the JWT configuration settings.
type JWTConfig struct {
	ISSUER     string
	EXP_MIN    string
	SECRET_KEY string
}

// AuthCtl is a struct that represents an authentication controller.
type AuthCtl struct {
	name      string
	apikey    string
	jwtConfig *JWTConfig
	dataStore map[string]DataPointExtended
}

var (
	ErrAuthFail                error = errors.New("[API-Users]: authentication failure")
	ErrNotAllowedRequestMethod error = errors.New("[API-Users]: requst method is not allowed for this endpoint")
	ErrUserExisted             error = errors.New("[API-Users]: user already existed")
	ErrEnvNotLoaded                  = errors.New("[JWT]: fail to load the env file")
	ErrPayloadParsing                = errors.New("[JWT]: fail to parse payload")
)

// NewAuthCtl sets:
// - the apikey to use to connect to SingPass API Service
// - the name assigned to this struct (for reference purpose)
func NewAuthCtl(name string, apikey string, jwtConfig *JWTConfig, ds map[string]DataPointExtended) *AuthCtl {
	return &AuthCtl{
		name:      name,
		apikey:    apikey,
		jwtConfig: jwtConfig,
		dataStore: ds,
	}
}

// Auth executes the authentication flow using SingPass API Service.
func (a *AuthCtl) Auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// reqBody struct to facilitate request body json unmarshalling.
	rb := struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}{}

	err := utils.ParseBody(r, &rb)
	if err != nil {
		customErr := errors.New(`[AUTH-CTL] fail to parse JSON body`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	dataPt := DataPointExtended{}

	// determin if in-memory cache or microservice will be used in the proceeding steps of the flow.
	if a.dataStore != nil {
		// in-memory cache is available.
		// authenticate with support from in-memory cache.

		// find the user data point, by phone number as the key.
		// found, err := a.dataStore.Find(rb.Phone)
		found := a.dataStore[rb.Phone]
		if found.UserID == 0 && found.Password == "" {
			customErr := errors.New(`[AUTH-CTL] fail to authenticate`)
			utils.SendErrorMsgToClient(&w, customErr)
			return
		}

		// ok.  found the data point.

		// ok. compare password and password hash (from data point) using bcrypt.
		err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(rb.Password))
		if err != nil {
			customErr := errors.New(`[AUTH-CTL] fail to authenticate`)
			utils.SendErrorMsgToClient(&w, customErr)
			return
		}

		// ok. passed all checks.
		// ready to generate JWT.
		dataPt = found
	}

	// in-memory cache is NOT available
	// authenticate with support from microservice.

	// set the jwt issuer value
	JWT_ISSUER := a.jwtConfig.ISSUER

	// // set the jwt expiry time lapse (in minutes)
	JWT_EXP_MINUTES, err := strconv.Atoi(a.jwtConfig.EXP_MIN)
	if err != nil {
		customErr := errors.New(`[AUTH-CTL] fail to set jwt expiry time frame`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	exp := time.Now().Add(time.Minute * time.Duration(JWT_EXP_MINUTES)).UnixMilli()
	pl := utils.JWTPayload{
		Id:    dataPt.UserID,
		Name:  dataPt.Name,
		Phone: dataPt.Phone,
		Iss:   JWT_ISSUER,
		Exp:   exp,
	}

	// ok.
	// generate JWT.
	token := ""
	token, err = generateJWT(pl, *a.jwtConfig)
	if err != nil {
		customErr := errors.New(`[AUTH-CTL] fail to generate JWT`)
		utils.SendForbiddenMsgToClient(&w, customErr)
		return
	}

	// set the response header attribute, "Authorization"
	bearerToken := fmt.Sprintf("Bearer %s", token)
	w.Header().Set("Authorization", bearerToken)

	// prepare the response body json
	respBody := fmt.Sprintf(`{
		"ok": true,
		"msg": "[AUTH-CTL]: authentication ok",
		"data": {
			"id" : %d,
			"token": "%s",
			"name" : "%s",
			"phone" : "%s"
		}
	}`, int64(dataPt.UserID), token, dataPt.Name, dataPt.Phone)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// Register add a user into the users microservice and cache the user data point returned by the microservice into the in-memory cache.
func (a *AuthCtl) Register(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// reqBody struct to facilitate INBOUND request body json unmarshalling.
	type postInput struct {
		Phone    string `json:"phone"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	rb := postInput{}
	err := utils.ParseBody(r, &rb)
	if err != nil {
		customErr := errors.New(`[AUTH-CTL] fail to parse JSON body`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// bcrypt the password.
	pwhash, err := bcrypt.GenerateFromPassword([]byte(rb.Password), bcrypt.MinCost)
	if err != nil {
		customErr := errors.New(`[AUTH-CTL] fail to process the password`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// dataPt := DataPoint{}
	dataPoints := []DataPoint{}

	if a.dataStore != nil {
		// in-memory cache is available.
		// register with support from in-memory cache.

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
		endpoint := fmt.Sprintf(`%s/adduser`, API_ROOT_URL)

		// POST request body json to be send to microservice.
		sendInput := postInput{
			Name:     rb.Name,
			Phone:    rb.Phone,
			Password: string(pwhash),
		}

		reqBody, err := json.Marshal(sendInput)
		if err != nil {
			customErr := errors.New(`[AUTH-CTL] fail to parse request body`)
			utils.SendErrorMsgToClient(&w, customErr)
			return
		}

		// prepare the POST request.
		apiReq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
		if err != nil {
			customErr := errors.New(`[AUTH-CTL] registration, failed`)
			utils.SendErrorMsgToClient(&w, customErr)
			return
		}

		// set all the required header attributes of this GET request.
		apiReq.Header.Set("Content-Type", "application/json")
		apiReq.Header.Set("apiKey", API_KEY)
		apiReq.Header.Set("username", API_USERNAME)

		// check if user already exist in microservice, by phone number.

		// is new user (based on phone number).
		// send out the POST request.
		outcome, err := client.Do(apiReq)
		if err != nil {
			customErr := errors.New(`[AUTH-CTL] registration, failed`)
			utils.SendErrorMsgToClient(&w, customErr)
			return
		}

		type respBody struct {
			Ok   bool                   `json:"ok"`
			Msg  string                 `json:"msg"`
			Data map[string]interface{} `json:"data"`
		}

		var outcomeRespBody respBody
		err = utils.ParseResponseBody(outcome, &outcomeRespBody)
		if err != nil {
			customErr := errors.New(`[AUTH-CTL] registration, failed`)
			utils.SendErrorMsgToClient(&w, customErr)
			return
		}

		if !outcomeRespBody.Ok && (outcomeRespBody.Msg == "[MS-Users]- Duplicate user.") {
			customErr := errors.New(outcomeRespBody.Msg)
			utils.SendBadRequestMsgToClient(&w, customErr)
			return
		}

		if !outcomeRespBody.Ok {
			customErr := errors.New(outcomeRespBody.Msg)
			utils.SendBadRequestMsgToClient(&w, customErr)
			return
		}

		// ok. breakdown the data.
		for _, data := range outcomeRespBody.Data {

			// break down the first level map.
			dp := DataPoint{}
			// loop through the second level map.
			// get the attribute and its value.
			// build the user data point struct.
			for k, v := range data.(map[string]interface{}) {

				switch k {
				case "UserID":
					dp.UserID = v.(float64)
				case "Phone":
					dp.Phone = v.(string)
				case "Name":
					dp.Name = v.(string)
				case "Points":
					dp.Points = v.(float64)
				case "LastLogin":
					dp.LastLogin = v.(string)
				}
				//
			}
			dataPoints = append(dataPoints, dp)

		}
	}

	// data in json string format
	data, err := json.Marshal(dataPoints)
	if err != nil {
		customErr := errors.New(`[AUTH-CTL] fail to parse result into json`)
		utils.SendErrorMsgToClient(&w, customErr)
		return
	}

	// ok.
	a.dataStore[dataPoints[0].Phone] = DataPointExtended{
		UserID:    dataPoints[0].UserID,
		Phone:     dataPoints[0].Phone,
		Name:      dataPoints[0].Name,
		Points:    dataPoints[0].Points,
		LastLogin: dataPoints[0].LastLogin,
		Password:  string(pwhash),
	}
	utils.SendDataToClient(&w, data, "registration, successful")
	//
}

// VerifyToken works in complement with the middleware that execute the token verification checks.
// When the request reach this endpint, it has a valid token.  This endpoint will return the verified JSON response body.
func (a *AuthCtl) VerifyToken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//
	authorization := r.Header.Get("Authorization")

	msg := fmt.Sprintln(`{
		"ok" : true,
		"msg" : "[AUTH-CTL]: token is valid",
		"data" : {}
	}`)

	w.Header().Set("Authorization", authorization)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
	//
}

// generateJWT will generate a JWT using the header and payload passed in.
func generateJWT(payload utils.JWTPayload, config JWTConfig) (string, error) {

	header := `{
		"alg": "SHA512",
		"typ" : "JWT"
	}`

	// convert payload data to json string
	pl, err := json.Marshal(payload)
	if err != nil {
		return "", ErrPayloadParsing
	}

	token := utils.Generate(header, string(pl), config.SECRET_KEY)

	return token, nil
}
