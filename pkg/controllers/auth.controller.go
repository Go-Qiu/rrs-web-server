package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	dataStore map[string]DataPoint
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
func NewAuthCtl(name string, apikey string, jwtConfig *JWTConfig, ds map[string]DataPoint) *AuthCtl {
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

	// dataPont struct for handling user data and instance assertion purpose.

	// userPt := dataPoint{}

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

	dataPt := DataPoint{}

	// determin if in-memory cache or microservice will be used in the proceeding steps of the flow.
	if a.dataStore != nil {
		// in-memory cache is available.
		// authenticate with support from in-memory cache.

		// find the user data point.
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
