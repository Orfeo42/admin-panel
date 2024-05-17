package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want *time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDate(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimePToForm(t *testing.T) {
	type args struct {
		valueFrom *time.Time
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
			if got := FormatTimePToForm(tt.args.valueFrom); got != tt.want {
				t.Errorf("FormatTimePToForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeToForm(t *testing.T) {
	type args struct {
		valueFrom time.Time
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
			if got := FormatTimeToForm(tt.args.valueFrom); got != tt.want {
				t.Errorf("FormatTimeToForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimePToTable(t *testing.T) {
	type args struct {
		valueFrom *time.Time
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
			if got := FormatTimePToTable(tt.args.valueFrom); got != tt.want {
				t.Errorf("FormatTimePToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeToTable(t *testing.T) {
	type args struct {
		valueFrom time.Time
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
			if got := FormatTimeToTable(tt.args.valueFrom); got != tt.want {
				t.Errorf("FormatTimeToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToTime(t *testing.T) {
	type args struct {
		valueFrom string
	}
	tests := []struct {
		name string
		args args
		want *time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToTime(tt.args.valueFrom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
