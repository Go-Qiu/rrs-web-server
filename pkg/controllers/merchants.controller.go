package controllers

import (
	"fmt"
	"net/http"
)

type MerchantsCtl struct {
	name   string
	apikey string
}

func NewMerchantsCtl(name string, apikey string) *MerchantsCtl {
	return &MerchantsCtl{
		name:   name,
		apikey: apikey,
	}
}

func (m *MerchantsCtl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-MERCHANTS]: retrieval of list of merchants, successful",
		"data" : [%s]
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// GetById will retrieve the specific merchant, identified by the id.
func (m *MerchantsCtl) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// code to retrieve the merchant data.

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-MERCHANTS]: retrieved merchant data, successful",
		"data" : [%s]
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// Create will add a new merchant data (optionally, with branches)
func (m *MerchantsCtl) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// code to create merchant data.

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-MERCHANTS]: created merchant data, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

//
func (m *MerchantsCtl) AddBranches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// code to add branches to a specific merchant data.

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-MERCHANTS]: branches added to merchant data, successful",
		"data" : [%s]
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

func (m *MerchantsCtl) RemoveBranches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// code to remove specific branches of a specific merchant data.

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-MERCHANTS]: branches removed for merchant data, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// Deactivate will de-activate a merchant data.
func (m *MerchantsCtl) HandlePutRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// action := "DEACTIVATE-MERCHANT"
	action := "ACTIVATE-MERCHANT"

	var respBody string

	if action == "DEACTIVATE-MERCHANT" {

		// code to deactivate merchant data

		// data in json string format
		data := ""

		respBody = fmt.Sprintf(`{
			"ok" : true,
			"msg" : "[MS-MERCHANTS]: deactivated merchant data, successful",
			"data" : [%s]
		}`, data)
		//
	} else if action == "ACTIVATE-MERCHANT" {

		// code to activate merchant data

		// data in json string format
		data := ""

		respBody = fmt.Sprintf(`{
			"ok" : true,
			"msg" : "[MS-MERCHANTS]: activated merchant data, successful",
			"data" : [%s]
		}`, data)
		//
	}

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}
