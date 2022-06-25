package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// TransactCtl is a struct that represents a user controller.
type VouchersCtl struct {
	name      string
	apikey    string
	dataStore map[string]DataPointExtended
}

func NewVouchersCtl(name string, apikey string, ds map[string]DataPointExtended) *VouchersCtl {
	return &VouchersCtl{
		name:      name,
		apikey:    apikey,
		dataStore: ds,
	}
}

// Create executes the flow to add a recyclable transaction and assocate it to a specific user.
func (t *VouchersCtl) Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// get the user id
	params := mux.Vars(r)
	id := params["id"]

	// data in json string format
	data := fmt.Sprintf(`{
		"user": {
			"id": %s
		}
	}`, id)

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-USERS]: added a transaction, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
}
