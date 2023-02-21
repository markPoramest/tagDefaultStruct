package main

import "testing"

func TestTagDefault(t *testing.T) {
	type TestStruct struct {
		StringField string  `default:"test"`
		BoolField   bool    `default:"true"`
		IntField    int     `default:"123"`
		FloatField  float64 `default:"3.14"`
	}

	testStruct := &TestStruct{}
	TagDefault(testStruct)

	expected := TestStruct{
		StringField: "test",
		BoolField:   true,
		IntField:    123,
		FloatField:  3.14,
	}

	if *testStruct != expected {
		t.Errorf("TagDefault returned %v, expected %v", *testStruct, expected)
	}
}
