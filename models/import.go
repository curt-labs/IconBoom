package models

import (
	"log"

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
	log.Printf("Begin testing %d vehicles", len(vs))
	return TryAllVehiclesInHeadlessBrowser(vs)
}
