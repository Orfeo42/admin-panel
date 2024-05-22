package utils

import (
	"reflect"
	"testing"
)

func TestParseString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatStringToForm(t *testing.T) {
	type args struct {
		valueFrom *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatStringToForm(tt.args.valueFrom); got != tt.want {
				t.Errorf("FormatStringToForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
