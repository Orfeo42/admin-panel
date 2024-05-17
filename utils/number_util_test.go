package utils

import (
	"reflect"
	"testing"
)

func TestStringToUint(t *testing.T) {
	type args struct {
		valueFrom string
	}
	tests := []struct {
		name    string
		args    args
		want    *uint
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToUint(tt.args.valueFrom)
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

func TestStringToInt(t *testing.T) {
	type args struct {
		valueFrom string
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt(tt.args.valueFrom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToString(t *testing.T) {
	type args struct {
		valueFrom string
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
			if got := StringToString(tt.args.valueFrom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToString() = %v, want %v", got, tt.want)
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

func TestFormatUintToFormString(t *testing.T) {
	type args struct {
		valueFrom *uint
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
			if got := FormatUintToFormString(tt.args.valueFrom); got != tt.want {
				t.Errorf("FormatUintToFormString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintToString(t *testing.T) {
	type args struct {
		valueFrom uint
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
			if got := UintToString(tt.args.valueFrom); got != tt.want {
				t.Errorf("UintToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_round(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := round(tt.args.num); got != tt.want {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFixed(tt.args.num, tt.args.precision); got != tt.want {
				t.Errorf("ToFixed() = %v, want %v", got, tt.want)
			}
		})
	}
}
