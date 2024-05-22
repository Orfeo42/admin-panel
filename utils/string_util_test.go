package utils

import (
	"reflect"
	"testing"
)

func TestParseString(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  *string
	}{
		{name: "enmpty string as input", value: "", want: nil},
		{name: "complete string as input", value: "ciao pippo", want: StringPtr("ciao pippo")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseString(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatStringToForm(t *testing.T) {
	tests := []struct {
		name      string
		valueFrom *string
		want      string
	}{
		{name: "nil as input", valueFrom: nil, want: ""},
		{name: "String pointer as input", valueFrom: StringPtr("Ciao"), want: "Ciao"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatStringToForm(tt.valueFrom); got != tt.want {
				t.Errorf("FormatStringToForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
