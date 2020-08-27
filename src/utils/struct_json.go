package utils

import "encoding/json"

func Struct2JsonStr(p interface{}) (string, error) {
	jsonBytes, err := json.Marshal(p)
	jsonstr := string(jsonBytes)
	return jsonstr, err
}

func JsonStr2Struct(jsonstr string, p interface{}) error {
	err := json.Unmarshal([]byte(jsonstr), p)
	return err
}
