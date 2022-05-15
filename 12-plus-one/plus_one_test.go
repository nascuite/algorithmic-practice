package main

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "hundred",
			args: args{[]int{1, 2, 3}},
			want: []int{1, 2, 4},
		},
		{
			name: "thousand",
			args: args{[]int{4, 3, 2, 1}},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "single",
			args: args{[]int{9}},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
