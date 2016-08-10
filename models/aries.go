package models

import (
	"fmt"
	"strings"

	"github.com/aries-auto/envision-api"
	"github.com/curt-labs/go-utensils/mongo"
	"gopkg.in/mgo.v2/bson"
)

// Creates a map of aries vehicles
// Use to check against envision vehicles
func GetAriesVehicles() (map[string]envisionAPI.Vehicle, error) {
	vmap := make(map[string]envisionAPI.Vehicle)
	vs, err := FromDb()
	if err != nil {
		return vmap, err
	}
	for _, v := range vs {
		key := fmt.Sprintf("%s:%s:%s", v.Year, v.Make, v.Model)
		vmap[key] = v
	}
	return vmap, nil
}

func FromDb() ([]envisionAPI.Vehicle, error) {
	var vs []envisionAPI.Vehicle
	err := mongo.InitMongo("product_data")
	if err != nil {
		return vs, err
	}
	query := bson.M{"brand.id": 3, "vehicle_applications.year": bson.M{"$exists": true}}
	result := []struct {
		Vehicle_Applications []struct {
			Year  string `bson:"year"`
			Make  string `bson:"make"`
			Model string `bson:"model"`
		} `bson:"vehicle_applications"`
	}{}
	err = mongo.MongoSession.DB("product_data").C("products").Find(query).Select(bson.M{"vehicle_applications.year": 1, "vehicle_applications.make": 1, "vehicle_applications.model": 1}).All(&result)
	for _, res := range result {
		for _, resv := range res.Vehicle_Applications {
			vs = append(vs, envisionAPI.Vehicle{Year: resv.Year, Make: resv.Make, Model: resv.Model})
			modelNameArray := strings.Split(resv.Model, " ")
			if len(modelNameArray) < 2 {
				continue
			}
			vs = append(vs, envisionAPI.Vehicle{Year: resv.Year, Make: resv.Make, Model: modelNameArray[0]}) // e.g. Silverado 1500 => Silverado
		}
	}
	return vs, err
}
