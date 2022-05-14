package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "III",
			args: args{"III"},
			want: 3,
		},
		{
			name: "LVIII",
			args: args{"LVIII"},
			want: 58,
		},
		{
			name: "MCMXCIV",
			args: args{"MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
