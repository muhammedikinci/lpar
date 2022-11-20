package lpar

import "fmt"

type Parameter map[string]interface{}

func (param Parameter) With(key string, value interface{}) Parameter {
	param[key] = value
	return param
}

func (param Parameter) String() string {
	parameters := ""

	for key, value := range param {
		parameters += fmt.Sprintf("%s: %v\n", key, value)
	}

	return parameters
}

func Param(key string, value interface{}) Parameter {
	return Parameter{
		key: value,
	}
}
