package main

import "testing"

func Test_singleNumberFast(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "last",
			args: args{[]int{2, 2, 1}},
			want: 1,
		},
		{
			name: "first",
			args: args{[]int{4, 1, 2, 1, 2}},
			want: 4,
		},
		{
			name: "one",
			args: args{[]int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberXOR(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumberSlow(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "last",
			args: args{[]int{2, 2, 1}},
			want: 1,
		},
		{
			name: "first",
			args: args{[]int{4, 1, 2, 1, 2}},
			want: 4,
		},
		{
			name: "one",
			args: args{[]int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
