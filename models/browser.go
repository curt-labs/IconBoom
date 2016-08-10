package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aries-auto/envision-api"
)

var (
	ICON_USER   = ""
	ICON_PASS   = ""
	ICON_DOMAIN = ""
)

func init() {
	if user := os.Getenv("ICON_USER"); user != "" {
		ICON_USER = user
	}
	if pass := os.Getenv("ICON_PASS"); pass != "" {
		ICON_PASS = pass
	}
	if domain := os.Getenv("ICON_DOMAIN"); domain != "" {
		ICON_DOMAIN = domain
	}
}

func TryAllVehiclesInHeadlessBrowser(vs []envisionAPI.Vehicle) error {
	var err error
	successMap := make(map[string]envisionAPI.Vehicle)
	errorMap := make(map[string]envisionAPI.Vehicle)

	for i, v := range vs {
		vs, err := GetVehicleIDs(v)
		if err != nil && strings.Contains(err.Error(), "invalid character ','") {
			continue
		}
		if err != nil {
			return err
		}
		for _, vehicle := range vs {

			legit, err := headlessBrowser(vehicle.ID)
			if err != nil {
				return err
			}
			if legit {
				successMap[vehicle.ID] = vehicle
			} else {
				errorMap[vehicle.ID] = vehicle
			}
		}
	}
	for _, vehicle := range successMap {
		logSuccess(vehicle)
	}
	for _, vehicle := range errorMap {
		logError(vehicle)
	}

	return err
}

func headlessBrowser(id string) (bool, error) {
	log.Print("V ID: ", id)
	api := fmt.Sprintf("http://www.iconfigurators.com/ap-json/get-AR-reference-image-color.aspx?vehicle=%s&color=0&uid=0", id)

	resp, err := http.Get(api)
	if err != nil {
		return false, err
	}
	apiRes := struct {
		Result int `json:"Result"`
		Img    []struct {
			Src   string `json:"src"`
			Wheel int    `json:"wheel"`
		} `json:"img"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&apiRes)
	if err != nil && strings.Contains(err.Error(), "invalid character 'V'") {
		err = nil
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, nil
}
