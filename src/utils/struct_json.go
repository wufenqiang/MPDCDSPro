package utils

import "encoding/json"

func StructToJsonStr(p interface{}) (string, error) {
	jsonBytes, err := json.Marshal(p)
	jsonstr := string(jsonBytes)
	return jsonstr, err
}

func JsonStrToStruct(jsonstr string, p interface{}) error {
	err := json.Unmarshal([]byte(jsonstr), p)
	return err
}
