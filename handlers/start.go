package handlers

import (
	"html/template"
	"net/http"

	"github.com/curt-labs/IconBoom/models"
)

func Start(w http.ResponseWriter, r *http.Request) {
	err := models.TryAllVehicles()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte("YEAH!"))
	return
}

func Test(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	vehicleID := r.URL.Query().Get("v")
	data["v"] = vehicleID

	t, err := template.ParseFiles("templates/icon.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	t.Execute(w, data)
	return
}
