package controllers

import (
	"errors"
	"fmt"
	"net/http"
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
func NewAuthCtl(name string, apikey string, jwtConfig *JWTConfig) *AuthCtl {
	return &AuthCtl{
		name:      name,
		apikey:    apikey,
		jwtConfig: jwtConfig,
	}
}

// Auth executes the authentication flow using SingPass API Service.
func (a *AuthCtl) Auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// set the jwt issuer value
	// JWT_ISSUER := a.jwtConfig.ISSUER

	// set the jwt expiry time lapse (in minutes)
	// JWT_EXP_MINUTES, err := strconv.Atoi(a.jwtConfig.EXP_MIN)
	// if err != nil {
	// 	customErr := errors.New(`[AUTH-CTL] fail to set jwt expiry time frame`)
	// 	utils.SendErrorMsgToClient(&w, customErr)
	// 	return
	// }

	// auth code here.

	// ok.
	// generate JWT.

	// exp := time.Now().Add(time.Minute * time.Duration(JWT_EXP_MINUTES)).UnixMilli()
	// pl := utils.JWTPayload{
	// 	Id:        found.Email,
	// 	NameFirst: found.NameFirst,
	// 	NameLast:  found.NameLast,
	// 	IsAgent:   found.IsAgent,
	// 	IsActive:  found.IsActive,
	// 	Iss:       JWT_ISSUER,
	// 	Exp:       exp,
	// }

	// to-do:
	// get jwt from SingPass (to be coded further)
	token := "singpass"

	// set the response header attribute, "Authorization"
	bearerToken := fmt.Sprintf("Bearer %s", token)
	w.Header().Set("Authorization", bearerToken)

	// prepare the response body json
	respBody := fmt.Sprintf(`{
		"ok": true,
		"msg": "[AUTH-CTL]: authentication ok",
		"data": {
			"token": "%s"
		}
	}`, token)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}
