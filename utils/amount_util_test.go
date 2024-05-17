package utils

import (
	"testing"
)

func TestParseAmount(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{name: "simple integer input", arg: "10", want: 1000},
		{name: "integer with decimal but 0", arg: "10,00", want: 1000},
		{name: "float input whith 2 decimal", arg: "12,99", want: 1299},
		{name: "float input whith 4 decimal", arg: "10002,3594", want: 1000235},
		{name: "string input", arg: "pippo", want: 0},
		{name: "empty string input", arg: "", want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ParseAmount(tc.arg); got != tc.want {
				t.Errorf("ParseAmount() = %d, want %d", got, tc.want)
			}
		})
	}
}

func Test_formatFloat(t *testing.T) {

	tests := []struct {
		name string
		arg  float64
		want string
	}{
		{name: "simple integer input", arg: 1000, want: "1000.00"},
		{name: "0 input", arg: 0, want: "0.00"},
		{name: "float input whith 2 decimal", arg: 10.22, want: "10.22"},
		{name: "float input whith 4 decimal", arg: 0.5974, want: "0.60"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := formatFloat(tc.arg); got != tc.want {
				t.Errorf("formatFloat() = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestFormatAmount(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want string
	}{
		{name: "0 as input", arg: 0, want: "0.00"},
		{name: "unit as input", arg: 1, want: "0.01"},
		{name: "ten as input", arg: 10, want: "0.10"},
		{name: "hundred as input", arg: 100, want: "1.00"},
		{name: "thousand as input", arg: 1000, want: "10.00"},
		{name: "ten thousand as input", arg: 10000, want: "100.00"},
		{name: "ten thousand with all number valued as input", arg: 29875, want: "298.75"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FormatAmount(tc.arg); got != tc.want {
				t.Errorf("FormatAmount() = %s, want %s", got, tc.want)
			}
		})
	}
}

func Test_formatIntWithSeparator(t *testing.T) {
	type args struct {
		integerPart string
		separator   string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0 as input", args: args{integerPart: "0", separator: "˙"}, want: "0"},
		{name: "ten as input", args: args{integerPart: "10", separator: "˙"}, want: "10"},
		{name: "hundred as input", args: args{integerPart: "100", separator: "˙"}, want: "100"},
		{name: "thousand as input", args: args{integerPart: "1000", separator: "˙"}, want: "1˙000"},
		{name: "ten thousand as input", args: args{integerPart: "10000", separator: "˙"}, want: "10˙000"},
		{name: "ten thousand with all number valued as input", args: args{integerPart: "29875", separator: "˙"}, want: "29˙875"},
		{name: "string input", args: args{integerPart: "pippo", separator: "˙"}, want: "pi˙ppo"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := formatIntWithSeparator(tc.args.integerPart, tc.args.separator); got != tc.want {
				t.Errorf("formatIntWithSeparator() = %s, want %s", got, tc.want)
			}
		})
	}
}

func Test_addThousandsSeparator(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want string
	}{
		{name: "0 as input", arg: "0", want: "0"},
		{name: "0 as input whith zero decimal", arg: "0.00", want: "0.00"},
		{name: "ten as input", arg: "10", want: "10"},
		{name: "hundred as input", arg: "100", want: "100"},
		{name: "thousand as input", arg: "1000", want: "1˙000"},
		{name: "ten thousand as input", arg: "10000", want: "10˙000"},
		{name: "ten thousand with all number valued as input", arg: "29875", want: "29˙875"},
		{name: "milion as input", arg: "9999999", want: "9˙999˙999"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := addThousandsSeparator(tc.arg); got != tc.want {
				t.Errorf("addThousandsSeparator() = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestAmountToString(t *testing.T) {

	tests := []struct {
		name string
		arg  float64
		want string
	}{
		{name: "float input whith 4 decimal", arg: 0.5974, want: "0.60"},
		{name: "float input whith 2 decimal", arg: 0.20, want: "0.20"},
		{name: "0 as input", arg: 0, want: "0.00"},
		{name: "10 as input", arg: 10, want: "10.00"},
		{name: "100 as input", arg: 100, want: "100.00"},
		{name: "1000 as input", arg: 1000, want: "1˙000.00"},
		{name: "10000 as input", arg: 10000, want: "10˙000.00"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := AmountToString(tc.arg); got != tc.want {
				t.Errorf("AmountToString() = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestAmountIntegerToString(t *testing.T) {

	tests := []struct {
		name string
		arg  int
		want string
	}{
		{name: "-1 as input", arg: -1, want: "-0.01"},
		{name: "0 as input", arg: 0, want: "0.00"},
		{name: "10 as input", arg: 10, want: "0.10"},
		{name: "100 as input", arg: 100, want: "1.00"},
		{name: "1000 as input", arg: 1000, want: "10.00"},
		{name: "10000 as input", arg: 10000, want: "100.00"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := AmountIntegerToString(tc.arg); got != tc.want {
				t.Errorf("AmountIntegerToString() = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestFormatIntToAmount(t *testing.T) {

	tests := []struct {
		name string
		arg  *int
		want string
	}{
		{name: "nil as input", arg: nil, want: ""},
		{name: "0 as input", arg: IntPtr(0), want: "0.00"},
		{name: "1 as input", arg: IntPtr(1), want: "0.01"},
		{name: "10 as input", arg: IntPtr(10), want: "0.10"},
		{name: "100 as input", arg: IntPtr(100), want: "1.00"},
		{name: "1000 as input", arg: IntPtr(1000), want: "10.00"},
		{name: "10000 as input", arg: IntPtr(10000), want: "100.00"},
		{name: "100000 as input", arg: IntPtr(100000), want: "1000.00"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FormatIntToAmount(tc.arg); got != tc.want {
				t.Errorf("AmountToString() = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestStringToAmount(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want *int
	}{
		{name: "0 as input", arg: "0", want: IntPtr(0)},
		{name: "1 as input", arg: "1", want: IntPtr(100)},
		{name: "input whith 4 decimal with 0 value", arg: "1,00", want: IntPtr(100)},
		{name: "input whith 2 decimal with value", arg: "1,22", want: IntPtr(122)},
		{name: "input whith 4 decimal", arg: "1,3594", want: IntPtr(135)},
		{name: "10 as input", arg: "10", want: IntPtr(1000)},
		{name: "100 as input", arg: "100", want: IntPtr(10000)},
		{name: "1000 as input", arg: "1000", want: IntPtr(100000)},
		{name: "10000 as input", arg: "10000", want: IntPtr(1000000)},
		{name: "string as input", arg: "pippo", want: nil},
		{name: "empty string as input", arg: "", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringToAmount(tt.arg)
			if got == nil {
				if got != tt.want {
					t.Errorf("StringToAmount() = %v, want %v", got, tt.want)
				}
				return
			}
			if *got != *tt.want {
				t.Errorf("StringToAmount() = %v, want %v", *got, *tt.want)
			}
		})
	}
}
