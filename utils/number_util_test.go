package utils

import (
	"reflect"
	"testing"
)

func TestStringToUintPtr(t *testing.T) {
	tests := []struct {
		name      string
		valueFrom string
		want      *uint
		wantErr   bool
	}{
		{name: "empty input", valueFrom: "", want: nil, wantErr: true},
		{name: "text as input", valueFrom: "pippo", want: nil, wantErr: true},
		{name: "number with decimal as input", valueFrom: "10.2", want: nil, wantErr: true},
		{name: "10 as input", valueFrom: "10", want: UintPtr(10), wantErr: false},
		{name: "100 as input", valueFrom: "100", want: UintPtr(100), wantErr: false},
		{name: "1000 as input", valueFrom: "1000", want: UintPtr(1000), wantErr: false},
		{name: "10000 as input", valueFrom: "10000", want: UintPtr(10000), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToUintPtr(tt.valueFrom)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintToString(t *testing.T) {
	tests := []struct {
		name      string
		valueFrom uint
		want      string
	}{
		{name: "0 as input", valueFrom: 0, want: "0"},
		{name: "1 as input", valueFrom: 1, want: "1"},
		{name: "10 as input", valueFrom: 10, want: "10"},
		{name: "100 as input", valueFrom: 100, want: "100"},
		{name: "1000 as input", valueFrom: 1000, want: "1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintToString(tt.valueFrom); got != tt.want {
				t.Errorf("UintToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_round(t *testing.T) {
	tests := []struct {
		name string
		num  float64
		want int
	}{
		{name: "0 as input", num: 0, want: 0},
		{name: "0.4 as input", num: 0.4, want: 0},
		{name: "0.5 as input", num: 0.5, want: 1},
		{name: "0.6 as input", num: 0.6, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := round(tt.num); got != tt.want {
				t.Errorf("round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFixed(t *testing.T) {
	type args struct {
		num       float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "0 as input", args: args{num: 0, precision: 2}, want: 0},
		{name: "10.30 as input", args: args{num: 10.30, precision: 2}, want: 10.3},
		{name: "10.344 as input", args: args{num: 10.344, precision: 2}, want: 10.34},
		{name: "10.345 as input", args: args{num: 10.345, precision: 2}, want: 10.35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFixed(tt.args.num, tt.args.precision); got != tt.want {
				t.Errorf("ToFixed() = %v, want %v", got, tt.want)
			}
		})
	}
}
