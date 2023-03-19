package main

import "testing"

func TestScanStringToBig(t *testing.T) {
	bigIntString := "29834082309480298304923804928300923804923"
	convertedValue := scanStringToBig(bigIntString)

	value, _ := convertedValue.MarshalText()
	if bigIntString != string(value) {
		t.Error("the value has not been correctly converted!")
	}
}
