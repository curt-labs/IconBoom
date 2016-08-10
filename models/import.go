package models

import (
	"github.com/aries-auto/envision-api"
)

func TryAllVehicles() error {
	AriesVehicleMap, err := GetAriesVehicles() //aries vehicles from Mongo DB
	if err != nil {
		return err
	}
	var vs []envisionAPI.Vehicle
	for _, ariesVehicle := range AriesVehicleMap {
		vs = append(vs, ariesVehicle)
	}
	return TryAllVehiclesInHeadlessBrowser(vs)
}
