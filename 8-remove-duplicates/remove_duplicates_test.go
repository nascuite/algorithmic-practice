package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums       []int
		ResultNums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{[]int{1, 1, 2}, []int{1, 2}},
			want: 2,
		},
		{
			name: "success",
			args: args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.nums[:got], tt.args.ResultNums) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums[:got], tt.args.ResultNums)
			}
		})
	}
}
