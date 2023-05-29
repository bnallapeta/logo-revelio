package store

import (
	"encoding/json"
	"os"
)

type LogoData struct {
	Logos map[string]string `json:"logos"`
}

var Logodata LogoData

func LoadLogosMap() error {
	// Read and load the logos mapping from JSON file
	logosFile, err := os.ReadFile("web/static/json/logos.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(logosFile, &Logodata)
	if err != nil {
		return err
	}

	return nil
}

func GetLogos() (map[string]string, error) {
	return Logodata.Logos, nil
}
