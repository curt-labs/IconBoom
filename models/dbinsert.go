package models

import (
	"encoding/csv"
	"fmt"
	"github.com/curt-labs/go-utensils/mongo"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Vehicle struct {
	Year    string
	Make    string
	Model   string
	IconIDs []int
}

// ImportSuccess inserts successfully envision-ed vehicles from the success.csv into an iconSuccess collection
func ImportSuccess() error {
	rows, err := importCsv("success.csv")
	if err != nil {
		return err
	}
	vehicleIDMap, err := compileVehicleMap(rows)
	if err != nil {
		return err
	}
	vehicles := aggregateVehicles(vehicleIDMap)

	err = insertInMongo(vehicles, "iconSuccess")
	return err
}

// ImportERrors inserts unsuccessfully envision-ed vehicles from the success.csv into an iconErrors collection
func ImportErrors() error {
	rows, err := importCsv("errors.csv")
	if err != nil {
		return err
	}
	vehicleIDMap, err := compileVehicleMap(rows)
	if err != nil {
		return err
	}
	vehicles := aggregateVehicles(vehicleIDMap)

	err = insertInMongo(vehicles, "iconErrors")
	return err
}

func importCsv(name string) ([][]string, error) {
	var rows [][]string
	f, err := os.Open(name)
	if err != nil {
		return rows, err
	}
	reader := csv.NewReader(f)
	return reader.ReadAll()
}

func compileVehicleMap(rows [][]string) (map[string][]int, error) {
	vehicleIDMap := make(map[string][]int)
	for _, row := range rows {
		if len(row) < 4 {
			continue
		}
		key := fmt.Sprintf("%s:%s:%s", row[1], row[2], row[3])
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return vehicleIDMap, err
		}
		if _, ok := vehicleIDMap[key]; !ok || !inIntArray(id, vehicleIDMap[key]) {
			vehicleIDMap[key] = append(vehicleIDMap[key], id)
		}
	}
	return vehicleIDMap, nil
}

func aggregateVehicles(vehicleIDMap map[string][]int) []Vehicle {
	var vehicles []Vehicle
	for key, ids := range vehicleIDMap {
		keyArray := strings.Split(key, ":")
		v := Vehicle{
			Year:    keyArray[0],
			Make:    keyArray[1],
			Model:   keyArray[2],
			IconIDs: ids,
		}
		vehicles = append(vehicles, v)
	}
	return vehicles
}

func inIntArray(i int, ints []int) bool {
	for _, in := range ints {
		if i == in {
			return true
		}
	}
	return false
}

func insertInMongo(vehicles []Vehicle, collection string) error {
	err := mongo.InitMongo("icon")
	if err != nil {
		return err
	}
	c := mongo.MongoSession.DB(mongo.MongoDatabase).C(collection)
	for _, vehicle := range vehicles {
		sort.Ints(vehicle.IconIDs)
		err = c.Insert(vehicle)
		if err != nil {
			return err
		}
	}
	return err
}
