package logging

import (
	"encoding/json"
	"github.com/h4lim/go-sdk/app/types"
	"io/ioutil"
)

const (
	M_ERROR_CODES_EN = "error_codes_en"
	M_ERROR_CODES_ID = "error_codes_id"
)

var maps = make(map[string](map[string]string))

type Map struct {
	Name    string            `json:"map_name"`
	Entries map[string]string `json:"entries"`
}

type MapsDocument struct {
	Maps []Map `json:"maps"`
}

func InitializeMaps(fileName string, localize string) *types.MessageMap {
	var messageMap types.MessageMap
	var mapName string
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf(INTERNAL, "Error reading maps file %s", err.Error())
	}

	if localize == "EN" {
		mapName = M_ERROR_CODES_EN
	} else {
		mapName = M_ERROR_CODES_ID
	}

	log.Debugf(INTERNAL, "found map data file map_data.json")
	doc := MapsDocument{}
	if err := json.Unmarshal(bytes, &doc); err != nil {
	}
	for _, elem := range doc.Maps {
		log.Debugf(INTERNAL, "populating map %s", elem.Name)
		if elem.Name == mapName {
			maps[elem.Name] = elem.Entries
			messageMap.Name = elem.Name
			messageMap.Entries = elem.Entries
		}
	}
	log.Debugf(INTERNAL, "parsed map data file map_data.json")

	return &messageMap
}
