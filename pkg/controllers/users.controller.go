package controllers

import (
	"fmt"
	"net/http"
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
