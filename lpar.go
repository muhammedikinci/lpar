package lpar

import (
	"bytes"
	"fmt"
)

type Parameter map[string]interface{}

func Param(key string, value interface{}) Parameter {
	return Parameter{
		key: value,
	}
}

func (param Parameter) With(key string, value interface{}) Parameter {
	param[key] = value
	return param
}

func (param Parameter) String() string {
	parameters := ""
	printTemplate := "%s: %v\n"

	setTemplate := func(key string, value interface{}) string {
		return fmt.Sprintf(printTemplate, key, value)
	}

	for key, value := range param {
		switch valueWithType := value.(type) {
		case bytes.Buffer:
			parameters += setTemplate(key, valueWithType.String())
		default:
			parameters += setTemplate(key, valueWithType)
		}
	}

	return parameters
}
