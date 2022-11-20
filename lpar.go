package lpar

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const (
	printTemplate    = "%s: %s\n"
	variableTemplate = "%v"
	jsonTemplate     = "\"%s\":\"%s\","
	comma            = ","
	leftCurlyBrace   = "{"
	rightCurlyBrace  = "}"
	newLine          = "\n"
	backQuote        = "\""
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

// Avoid implement string interface
func (param Parameter) AsString() string {
	parameters := ""

	for key, value := range param {
		parameters += setPrintTemplate(key, toString(value))
	}

	return strings.TrimRight(parameters, newLine)
}

func (param Parameter) AsJsonString() string {
	jsonString := leftCurlyBrace

	for key, value := range param {
		jsonString += setJsonTemplate(key, toString(value))
	}

	return strings.TrimRight(jsonString, comma) + rightCurlyBrace
}

func toString(value interface{}) string {
	valueString := ""

	kind := reflect.TypeOf(value).Kind()

	if kind == reflect.Struct {
		typeOf := reflect.TypeOf(value)

		for i := 0; i < reflect.TypeOf(value).NumField(); i++ {
			if typeOf.Field(i).Tag.Get("json") != "" {
				bytes, err := json.Marshal(value)
				if err != nil {
					return ""
				}

				valueString = fmt.Sprintf("%q", string(bytes))
				valueString = strings.Trim(valueString, backQuote)
				break
			}
		}
	} else {
		valueString = setVariableTemplate(value)
	}

	return valueString
}

func setPrintTemplate(key, value string) string {
	return fmt.Sprintf(printTemplate, key, value)
}

func setVariableTemplate(value interface{}) string {
	return fmt.Sprintf(variableTemplate, value)
}

func setJsonTemplate(key, value string) string {
	return fmt.Sprintf(jsonTemplate, key, value)
}
