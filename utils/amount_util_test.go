package utils

import (
	"testing"
)

func Test_ParseAmount(t *testing.T) {
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

func Test_FormatAmount(t *testing.T) {
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
		{name: "string input", args: args{integerPart: "pippo", separator: "˙"}, want: "pi˙ppo"},
		{name: "ten thousand with all number valued as input", args: args{integerPart: "29875", separator: "˙"}, want: "29˙875"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := formatIntWithSeparator(tc.args.integerPart, tc.args.separator); got != tc.want {
				t.Errorf("formatIntWithSeparator() = %s, want %s", got, tc.want)
			}
		})
	}
}

func Test_stringAmountToString(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want string
	}{
		{name: "", arg: "0", want: "0"},
		{name: "", arg: "0.00", want: "0.00"},
		{name: "", arg: "10", want: "10"},
		{name: "", arg: "100", want: "100"},
		{name: "", arg: "1000", want: "1˙000"},
		{name: "", arg: "10000", want: "10˙000"},
		{name: "", arg: "29875", want: "29˙875"},
		{name: "", arg: "99999999", want: "99˙999˙999"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := stringAmountToString(tc.arg); got != tc.want {
				t.Errorf("stringAmountToString() = %s, want %s", got, tc.want)
			}
		})
	}
}

func Test_AmountToString(t *testing.T) {

	tests := []struct {
		name string
		arg  float64
		want string
	}{
		{name: "", arg: 1000, want: "1˙000.00"},
		{name: "", arg: 10, want: "10.00"},
		{name: "", arg: 0, want: "0.00"},
		{name: "", arg: 0.5974, want: "0.60"},
		{name: "", arg: 9990.5974, want: "9˙990.60"},
		{name: "", arg: 999999990, want: "999˙999˙990.00"},
		{name: "", arg: 9990, want: "9˙990.00"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := AmountToString(tc.arg); got != tc.want {
				t.Errorf("AmountToString() = %s, want %s", got, tc.want)
			}
		})
	}
}

func Test_AmountIntegerToString(t *testing.T) {

	tests := []struct {
		name string
		arg  int
		want string
	}{
		{name: "", arg: 1000, want: "10.00"},
		{name: "", arg: 10, want: "0.10"},
		{name: "", arg: 0, want: "0.00"},
		{name: "", arg: 999999990, want: "9˙999˙999.90"},
		{name: "", arg: 9990, want: "99.90"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := AmountIntegerToString(tc.arg); got != tc.want {
				t.Errorf("AmountToString() = %s, want %s", got, tc.want)
			}
		})
	}
}

func Test_FormatIntToAmount(t *testing.T) {

	tests := []struct {
		name string
		arg  *int
		want string
	}{
		{name: "", arg: nil, want: ""},
		{name: "", arg: IntPtr(10), want: "0.10"},
		{name: "", arg: IntPtr(0), want: "0.00"},
		{name: "", arg: IntPtr(999999990), want: "9999999.90"},
		{name: "", arg: IntPtr(9990), want: "99.90"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FormatIntToAmount(tc.arg); got != tc.want {
				t.Errorf("AmountToString() = %s, want %s", got, tc.want)
			}
		})
	}
}

func Test_StringToAmount(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want *int
	}{
		{name: "", arg: "10", want: IntPtr(1000)},
		{name: "", arg: "10,00", want: IntPtr(1000)},
		{name: "", arg: "12", want: IntPtr(1200)},
		{name: "", arg: "12,00", want: IntPtr(1200)},
		{name: "", arg: "12,99", want: IntPtr(1299)},
		{name: "", arg: "10002,3594", want: IntPtr(1000235)},
		{name: "", arg: "1359782", want: IntPtr(135978200)},
		{name: "", arg: "987456321", want: IntPtr(98745632100)},
		{name: "", arg: "pippo", want: nil},
		{name: "", arg: "", want: nil},
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
				t.Errorf("StringToAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
