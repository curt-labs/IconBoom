package models

import (
	"testing"

	"github.com/aries-auto/envision-api"
)

func TestVehicle(t *testing.T) {
	vs := []envisionAPI.Vehicle{{
		Year:  "2015",
		Make:  "Toyota",
		Model: "Tacoma",
	}}
	err := TryAllVehiclesInHeadlessBrowser(vs)
	if err != nil {
		t.Log(err)
	}
}

func TestBrowser(t *testing.T) {
	ok, err := headlessBrowser("2040")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Log("Expected image to be not found")
	}

	ok, err = headlessBrowser("1375")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Log("Expected image to be  found")
	}
}
