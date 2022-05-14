package main

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "success-last-letter",
			args: args{"abc", "c"},
			want: 2,
		},
		{
			name: "success-letter-in-the-middle",
			args: args{"hello", "ll"},
			want: 2,
		},
		{
			name: "empty needle",
			args: args{"hello", ""},
			want: 0,
		},
		{
			name: "substring missing",
			args: args{"aaaaa", "bba"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
