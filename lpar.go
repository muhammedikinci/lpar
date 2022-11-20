package lpar

import (
	"bytes"
	"fmt"
	"strings"
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

const (
	printTemplate   = "%s: %s\n"
	stringTemplate  = "%v"
	jsonTemplate    = "\"%s\":\"%s\","
	comma           = ","
	leftCurlyBrace  = "{"
	rightCurlyBrace = "}"
)

func setPrintTemplate(key, value string) string {
	return fmt.Sprintf(printTemplate, key, value)
}

func setStringTemplate(value interface{}) string {
	return fmt.Sprintf(stringTemplate, value)
}

func setJsonTemplate(key, value string) string {
	return fmt.Sprintf(jsonTemplate, key, value)
}

func (param Parameter) String() string {
	parameters := ""

	for key, value := range param {
		parameters += setPrintTemplate(key, toString(value))
	}

	return parameters
}

func (param Parameter) Json() string {
	jsonString := leftCurlyBrace

	for key, value := range param {
		jsonString += setJsonTemplate(key, toString(value))
	}

	return strings.TrimRight(jsonString, comma) + rightCurlyBrace
}

func toString(value interface{}) string {
	valueString := ""

	switch valueWithType := value.(type) {
	case bytes.Buffer:
		valueString = valueWithType.String()
	default:
		valueString = setStringTemplate(valueWithType)
	}

	return valueString
}
