package controllers

import "net/http"

// CRUD interface for controller that needs to implemt CRUD features.
type CRUD interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	UpdateById(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
}

type CRUDController struct {
	name      string
	apikey    string
	dataStore map[string]DataPointExtended
	CRUD
}
