package utils

import (
	"bytes"
	"encoding/gob"
)

//RegGob *
func RegGob(o ...interface{}) {
	for _, v := range o {
		gob.Register(v)
	}
	mapGob := make(map[string]interface{})
	gob.Register(mapGob)
}

//GobMarshal *
func GobMarshal(v interface{}) ([]byte, error) {
	b := new(bytes.Buffer)
	err := gob.NewEncoder(b).Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

//GobUnmarshal *
func GobUnmarshal(data []byte, v interface{}) error {
	b := bytes.NewBuffer(data)
	return gob.NewDecoder(b).Decode(v)
}
