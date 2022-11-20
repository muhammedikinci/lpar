package lpar

type Parameter map[string]interface{}

func (param Parameter) With(key string, value interface{}) Parameter {
	param[key] = value
	return param
}

func Param(key string, value interface{}) Parameter {
	return Parameter{
		key: value,
	}
}
