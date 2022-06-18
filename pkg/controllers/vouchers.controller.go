package controllers

import (
	"fmt"
	"log"
	"net/http"
)

type VouchersCtl struct {
	name   string
	apikey string
}

type qty struct {
	dollars10 int
	dollars5  int
	dollars2  int
	dollars1  int
}
type InConvertPointsToVouchers struct {
	userId   string
	points   int
	vouchers struct {
		qty qty
	}
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

// GetByIdentifier retrieves a specific voucher data that matches the unique data id assigned at its creation.
func (v *VouchersCtl) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-VOUCHERS]: retrieval voucher data, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// RedeemById updates a specific voucher data (identified by Id) and change its status to 'REDEEMED'.
func (v *VouchersCtl) RedeemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-VOUCHERS]: redeemed voucher data, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

func (v *VouchersCtl) HandlePostRequest(w http.ResponseWriter, r *http.Request) {

	// need to replace with code that reads the action value from the request body.

	action := "CONVERT-POINTS-TO-VOUCHERS"
	inputs := InConvertPointsToVouchers{}
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	if action == "CONVET-POINT-TO-VOUCHERS" {
		err := convertPointsToVouchers(inputs)
		if err != nil {
			// error handling code (to be done)
			log.Println(err)
		}
	}

	// ok.
	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-VOUCHERS]: converted points to vouchers, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// convertPointsToVouchers will convert a specific number of reward points, for a specific user id, to a specific set of vouchers (different value quantum).
func convertPointsToVouchers(in InConvertPointsToVouchers) error {

	// call api endpoint at vouher system (to be done)

	return nil
}
