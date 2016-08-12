package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

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

	for _, v := range vs {
		vs, err := GetVehicleIDs(v)
		// if err != nil && strings.Contains(err.Error(), "invalid character ','") {
		// 	continue // No IconMedia vehicle for that year make model
		// }
		if err != nil {
			continue // No IconMedia vehicle for that year make model
		}
		for _, vehicle := range vs {

			legit, err := headlessBrowser(vehicle.ID)
			if legit && err == nil {
				if _, ok := successMap[vehicle.ID]; ok {
					continue //already logged
				}
				successMap[vehicle.ID] = vehicle
				logSuccess(vehicle, err)
			} else {
				if _, ok := errorMap[vehicle.ID]; ok {
					continue //already logged
				}
				errorMap[vehicle.ID] = vehicle
				logError(vehicle, err)
			}
			time.Sleep(time.Second * 2) //make it take forever
		}
	}
	return err
}

// Test the IconMedia endpoint that yields vehicle image
func headlessBrowser(id string) (bool, error) {
	log.Print("V ID: ", id)
	api := fmt.Sprintf("http://www.iconfigurators.com/ap-json/get-AR-reference-image-color.aspx?vehicle=%s&color=0&uid=0", id)
	resp, err := http.Get(api)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
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
		ok, err := tryVehicleImage(b)
		if err != nil {
			return false, err
		}
		return ok, nil
	}
	return false, nil
}

// iconMedia provides image in response, but does it really exist?
func tryVehicleImage(body []byte) (bool, error) {
	var err error
	strArray := strings.Split(string(body), "\"")
	var url string
	for _, str := range strArray {
		if strings.Contains(str, "http://images.iconfigurators.com") {
			url = str
			break
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, errors.New(fmt.Sprintf("%s does not yield a vehicle image successfully", url))
	}

	return true, nil
}
