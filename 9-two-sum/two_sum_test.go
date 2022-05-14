package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		wantA []int
		wantD []int
	}{
		// TODO: Add test cases.
		{
			name:  "success start",
			args:  args{[]int{2, 7, 11, 15}, 9},
			wantA: []int{0, 1},
			wantD: []int{1, 0},
		},
		{
			name:  "success last",
			args:  args{[]int{3, 2, 4}, 6},
			wantA: []int{1, 2},
			wantD: []int{2, 1},
		},
		{
			name:  "success short",
			args:  args{[]int{3, 3}, 6},
			wantA: []int{0, 1},
			wantD: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			if !(reflect.DeepEqual(got, tt.wantA) || reflect.DeepEqual(got, tt.wantD)) {
				t.Errorf("twoSum() = %v, want %v or %v", got, tt.wantA, tt.wantD)
			}
		})
	}
}
