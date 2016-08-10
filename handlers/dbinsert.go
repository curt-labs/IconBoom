package handlers

import (
	"net/http"

	"github.com/curt-labs/IconBoom/models"
)

// Insert Successfully Envision-ed vehicles into the iconSuccess collection
func InsertSuccess(w http.ResponseWriter, r *http.Request) {
	err := models.ImportSuccess()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Done w/ successes"))
	return
}

// Insert Unsuccessfully Envision-ed vehicles into the iconErrors collection
func InsertErrors(w http.ResponseWriter, r *http.Request) {
	err := models.ImportErrors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Done w/ errors"))
	return
}
