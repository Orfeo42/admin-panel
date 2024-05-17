package utils

import (
	"reflect"
	"testing"
)

func TestIntPtr(t *testing.T) {
	type args struct {
		i int
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
			if got := IntPtr(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
