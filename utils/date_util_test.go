package utils

import (
	"testing"
	"time"
)

func TestFormatTimePtrToForm(t *testing.T) {
	tests := []struct {
		name      string
		valueFrom *time.Time
		want      string
	}{
		{name: "nil input", valueFrom: nil, want: ""},
		{name: "valid Date as input", valueFrom: TimePtr(time.Date(1987, 05, 24, 0, 0, 0, 0, time.UTC)), want: "1987-05-24"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimePtrToForm(tt.valueFrom); got != tt.want {
				t.Errorf("FormatTimePToForm() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestFormatTimeToForm(t *testing.T) {
	tests := []struct {
		name      string
		valueFrom time.Time
		want      string
	}{
		{name: "valid Date as input", valueFrom: time.Date(1987, 05, 24, 0, 0, 0, 0, time.UTC), want: "1987-05-24"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimeToForm(tt.valueFrom); got != tt.want {
				t.Errorf("FormatTimeToForm() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestFormatTimePtrToTable(t *testing.T) {
	tests := []struct {
		name      string
		valueFrom *time.Time
		want      string
	}{
		{name: "nil input", valueFrom: nil, want: ""},
		{name: "valid Date as input", valueFrom: TimePtr(time.Date(1987, 05, 24, 0, 0, 0, 0, time.UTC)), want: "24-05-1987"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimePtrToTable(tt.valueFrom); got != tt.want {
				t.Errorf("FormatTimePToTable() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestFormatTimeToTable(t *testing.T) {
	tests := []struct {
		name      string
		valueFrom time.Time
		want      string
	}{
		{name: "valid Date as input", valueFrom: time.Date(1987, 05, 24, 0, 0, 0, 0, time.UTC), want: "24-05-1987"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimeToTable(tt.valueFrom); got != tt.want {
				t.Errorf("FormatTimeToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*func TestStringToTimePtr(t *testing.T) {
	tests := []struct {
		name      string
		valueFrom string
		want      *time.Time
	}{
		{name: "enmpty string as input", valueFrom: "", want: nil},
		{name: "wrong date as input", valueFrom: "1987-15-24", want: nil},
		{name: "valid Date as input", valueFrom: "1987-05-24", want: TimePtr(time.Date(1987, 05, 24, 0, 0, 0, 0, time.UTC))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToTimePtr(tt.valueFrom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}*/
