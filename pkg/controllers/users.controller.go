package controllers

import (
	"fmt"
	"net/http"
)

type UsersCtl struct {
	name   string
	apikey string
}

func NewUsersCtl(name string, apikey string) *UsersCtl {
	return &UsersCtl{
		name:   name,
		apikey: apikey,
	}
}

// GetAll retrieves all users data in the specified page index, given the number of records per page setting.
func (u *UsersCtl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-USERS]: retrieval of list of users, successful",
		"data" : [%s]
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// GetByIdentifier retrieves a specific user data that matches the unique data id assigned at its creation.
func (u *UsersCtl) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-USERS]: retrieval users data, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// Create will add a user data.  It will returns the created user data, retrieved from the database as a confirmation.
func (u *UsersCtl) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-USERS]: created users data, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// Delete will delete a specific user data and its associated transactions, vouchers and points.
func (u *UsersCtl) Inactivate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-USERS]: inactivated user data, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// AddTransaction will add a recycle transaction data to a specific user data.
func (u *UsersCtl) AddTransaction(w http.ResponseWriter, r *http.Request) {
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
	//
}

// XChangePointsToVouchers will pass in the specific number of reward points to be converted to a specific number of vouchers (of a specific monetary value).
func (u *UsersCtl) XChangePointsToVouchers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-USERS]: exchanged points to vouchers, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}

// RedeemVouchers will update the status of the specific vouchers in the list to 'REDEMED'.
func (u *UsersCtl) RedeemVouchers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// data in json string format
	data := ""

	respBody := fmt.Sprintf(`{
		"ok" : true,
		"msg" : "[MS-USERS]: redeemed vouchers, successful",
		"data" : {%s}
	}`, data)

	// send response to the requesting device
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respBody))
	//
}
