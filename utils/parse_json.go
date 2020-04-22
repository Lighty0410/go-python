package utils

import (
	"encoding/json"
	"io/ioutil"
)

// CreateJSONMap fills map from a json-file.
func CreateJSONMap(path string) (map[string]map[string]int, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]map[string]int)

	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		return nil, err
	}

	return jsonMap, nil
}
