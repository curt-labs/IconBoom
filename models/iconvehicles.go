package models

import (
	"github.com/aries-auto/envision-api"
)

func GetVehicleIDs(v envisionAPI.Vehicle) ([]envisionAPI.Vehicle, error) {
	var vehicles []envisionAPI.Vehicle

	conf, err := envisionAPI.NewConfig(ICON_USER, ICON_PASS, ICON_DOMAIN)
	if err != nil {
		return vehicles, err
	}

	vehicleResp, err := envisionAPI.GetVehicleByYearMakeModel(*conf, v.Year, v.Make, v.Model)
	if err != nil {
		return vehicles, err
	}

	for _, v := range vehicleResp.Vehicles {
		vehicles = append(vehicles, v)
	}
	return vehicles, err
}
