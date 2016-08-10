package models

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/aries-auto/envision-api"
)

func logError(v envisionAPI.Vehicle) error {
	fileName := "errors.csv"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	record := []string{strings.ToLower(strings.TrimSpace(v.ID)), strings.ToLower(strings.TrimSpace(v.Year)), strings.ToLower(strings.TrimSpace(v.Make)), strings.ToLower(strings.TrimSpace(v.Model)), strings.ToLower(strings.TrimSpace(v.BodyType))}
	err = writer.Write(record)
	writer.Flush()
	return err
}

func logSuccess(v envisionAPI.Vehicle) error {
	fileName := "success.csv"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	record := []string{strings.ToLower(strings.TrimSpace(v.ID)), strings.ToLower(strings.TrimSpace(v.Year)), strings.ToLower(strings.TrimSpace(v.Make)), strings.ToLower(strings.TrimSpace(v.Model)), strings.ToLower(strings.TrimSpace(v.BodyType))}
	err = writer.Write(record)
	writer.Flush()
	return err
}
