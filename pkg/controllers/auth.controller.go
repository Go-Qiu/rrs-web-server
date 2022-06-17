package controllers

import (
	"errors"
	"fmt"
	"net/http"
)

type AuthCtl struct {
	name   string
	apikey string
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
func NewAuthCtl(name string, apikey string) *AuthCtl {
	return &AuthCtl{
		name:   name,
		apikey: apikey,
	}
}

// Auth executes the authentication flow using SingPass API Service.
func (a *AuthCtl) Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
