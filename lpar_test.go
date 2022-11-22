package lpar

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParamWhenGettingNilValueShouldSaveValidValue(t *testing.T) {
	parameter := Param("test", nil)

	assert.Equal(t, nilValue, parameter["test"])
}

func TestAsStringWhenWorkingWithNilValueShouldReturnValidValue(t *testing.T) {
	parameter := Param("test", nil).AsString()

	assert.Equal(t, "test: nil", parameter)
}

func TestAsJsonStringWhenWorkingWithNilValueShouldReturnValidValue(t *testing.T) {
	parameter := Param("test", nil).AsJsonString()

	assert.Equal(t, "{\"test\": \"nil\"}", parameter)
}

func TestParamWhenGettingNumericValueShouldSaveValidValue(t *testing.T) {
	parameter := Param("test", 123)

	assert.Equal(t, 123, parameter["test"])
}

func TestAsStringWhenGettingByteBufferValueShouldReturnValidValue(t *testing.T) {
	bf := bytes.Buffer{}
	bf.WriteString("buffer string")

	parameter := Param("test", bf).AsString()

	assert.Equal(t, "test: buffer string", parameter)
}

func TestAsStringWhenGettingJsonableStructValueShouldReturnValidValue(t *testing.T) {
	type testStruct struct {
		JsonValue string `json:"json_value"`
	}

	jsonableConcrete := testStruct{JsonValue: "test"}

	parameter := Param("json_struct", jsonableConcrete).AsString()

	assert.Equal(t, "json_struct: {\\\"json_value\\\":\\\"test\\\"}", parameter)
}

func TestAsJsonStringWhenGettingJsonableStructValueShouldReturnValidValue(t *testing.T) {
	type testStruct struct {
		JsonValue string `json:"json_value"`
	}

	jsonableConcrete := testStruct{JsonValue: "test"}

	parameter := Param("json_struct", jsonableConcrete).AsJsonString()

	assert.Equal(t, "{\"json_struct\": \"{\\\"json_value\\\":\\\"test\\\"}\"}", parameter)
}

func TestAsJsonStringWhenGettingNormalStructWithoutAnyJsonTagValueShouldReturnValidValue(t *testing.T) {
	type testStruct struct {
		JsonValue string
	}

	jsonableConcrete := testStruct{JsonValue: "test"}

	parameter := Param("json_struct", jsonableConcrete).AsJsonString()

	assert.Equal(t, "{\"json_struct\": \"\"}", parameter)
}

func TestAsJsonStringWhenGettingNotValidJsonableStructValueShouldReturnValidValue(t *testing.T) {
	type testStruct struct {
		JsonValue *string `json:"json_value"`
	}

	jsonableConcrete := testStruct{JsonValue: nil}

	parameter := Param("json_struct", jsonableConcrete).AsJsonString()

	assert.Equal(t, "{\"json_struct\": \"{\\\"json_value\\\":null}\"}", parameter)
}

func TestAsStringWhenGettingMultipleDifferentValueShouldReturnValidValue(t *testing.T) {
	bf := bytes.Buffer{}
	bf.WriteString("buffer string")

	parameter := Param("test", bf).With("test2", "multiple").With("test3", 123).AsString()

	assert.Equal(t, "test: buffer string\ntest2: multiple\ntest3: 123", parameter)
}
