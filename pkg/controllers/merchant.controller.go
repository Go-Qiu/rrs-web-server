package controllers

import (
	"fmt"
	"net/http"
)

// UserCtl is a struct that represents a user controller.
type MerchantCtl struct {
	name      string
	apikey    string
	dataStore map[string]DataPointExtended
	CRUD
}

func NewMerchantCtl(name string, apikey string, ds map[string]DataPointExtended) *MerchantCtl {
	return &MerchantCtl{
		name:      name,
		apikey:    apikey,
		dataStore: ds,
	}
}

// AddTransaction executes the flow to add a recyclable transaction and assocate it to a specific user.
func (u *MerchantCtl) Create(w http.ResponseWriter, r *http.Request) {

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
