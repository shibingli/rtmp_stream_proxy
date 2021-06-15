package utils

import (
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

var jitJSON = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(v interface{}, jit ...bool) ([]byte, error) {
	if len(jit) > 0 && jit[0] {
		return jitJSON.Marshal(v)
	}

	return json.Marshal(v)
}

func JsonMarshalIndent(v interface{}, prefix, indent string, jit ...bool) ([]byte, error) {
	if len(jit) > 0 && jit[0] {
		return jitJSON.MarshalIndent(v, prefix, indent)
	}
	return json.MarshalIndent(v, prefix, indent)
}

func JsonMarshalToString(v interface{}, jit ...bool) (string, error) {
	if len(jit) > 0 && jit[0] {
		return jitJSON.MarshalToString(v)
	}

	b, err := json.Marshal(v)
	if nil != err {
		return "", err
	}

	return string(b), nil
}

func JsonUnmarshal(data []byte, v interface{}, jit ...bool) error {
	if len(jit) > 0 && jit[0] {
		return jitJSON.Unmarshal(data, v)
	}
	return json.Unmarshal(data, v)
}
