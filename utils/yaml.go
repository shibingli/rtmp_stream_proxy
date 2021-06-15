package utils

import (
	"gopkg.in/yaml.v3"
)

func YamlMarshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func YamlMarshalToString(v interface{}) (string, error) {
	b, err := YamlMarshal(v)
	if nil != err {
		return "", err
	}
	return string(b), nil
}

func YamlUnmarshal(data []byte, v interface{}, jit ...bool) error {
	return yaml.Unmarshal(data, v)
}

func YamlUnmarshalToString(data string, v interface{}, jit ...bool) error {
	return YamlUnmarshal([]byte(data), v)
}
