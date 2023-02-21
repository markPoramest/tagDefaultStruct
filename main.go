package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Test struct {
	Text        string  `default:"default text"`
	Number      int32   `default:"24"`
	FloatNumber float32 `default:"12.12"`
	Bool        bool    `default:"true"`
}

func main() {
	test := &Test{}
	TagDefault(test)
	fmt.Print(test)
}

func TagDefault(in interface{}) {
	value := reflect.ValueOf(in).Elem()

	parserData := mapperParserData()
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		defaultTag := field.Tag.Get("default")
		if defaultTag == "" {
			continue
		}
		fieldValue := value.Field(i)
		if fieldValue.CanSet() {
			parseDataFunc := parserData[fieldValue.Kind()]
			parseDataFunc(fieldValue, defaultTag)
		}
	}
}

func mapperParserData() map[reflect.Kind]func(reflect.Value, string) {
	return map[reflect.Kind]func(reflect.Value, string){
		reflect.String: func(fieldValue reflect.Value, defaultTag string) {
			fieldValue.SetString(defaultTag)
		},
		reflect.Int: func(fieldValue reflect.Value, defaultTag string) {
			defaultInt, err := strconv.ParseInt(defaultTag, 10, 64)
			if err == nil {
				fieldValue.SetInt(defaultInt)
			}
		},
		reflect.Int8: func(fieldValue reflect.Value, defaultTag string) {
			defaultInt, err := strconv.ParseInt(defaultTag, 10, 8)
			if err == nil {
				fieldValue.SetInt(defaultInt)
			}
		},
		reflect.Int16: func(fieldValue reflect.Value, defaultTag string) {
			defaultInt, err := strconv.ParseInt(defaultTag, 10, 16)
			if err == nil {
				fieldValue.SetInt(defaultInt)
			}
		},
		reflect.Int32: func(fieldValue reflect.Value, defaultTag string) {
			defaultInt, err := strconv.ParseInt(defaultTag, 10, 32)
			if err == nil {
				fieldValue.SetInt(defaultInt)
			}
		},
		reflect.Int64: func(fieldValue reflect.Value, defaultTag string) {
			defaultInt, err := strconv.ParseInt(defaultTag, 10, 64)
			if err == nil {
				fieldValue.SetInt(defaultInt)
			}
		},
		reflect.Float32: func(fieldValue reflect.Value, defaultTag string) {
			defaultFloat, err := strconv.ParseFloat(defaultTag, 32)
			if err == nil {
				fieldValue.SetFloat(defaultFloat)
			}
		},
		reflect.Float64: func(fieldValue reflect.Value, defaultTag string) {
			defaultFloat, err := strconv.ParseFloat(defaultTag, 64)
			if err == nil {
				fieldValue.SetFloat(defaultFloat)
			}
		},
		reflect.Bool: func(fieldValue reflect.Value, defaultTag string) {
			defaultBool, err := strconv.ParseBool(defaultTag)
			if err == nil {
				fieldValue.SetBool(defaultBool)
			}
		},
	}
}
