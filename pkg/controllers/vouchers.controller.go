package controllers

import (
	"fmt"
	"net/http"
)

type VouchersCtl struct {
	name   string
	apikey string
}

func NewVouchersCtl(name string, apikey string) *VouchersCtl {
	return &VouchersCtl{
		name:   name,
		apikey: apikey,
	}
}

func (v *VouchersCtl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-VOUCHERS]: retrieval of list of vouchers, successful",
		"data" : [%s]
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}
