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
