package main

import (
	"reflect"
	"testing"
)

// BenchmarkSimplestNTimes выполняет бенчмарк N раз.
func BenchmarkSimplestNTimesFast(b *testing.B) {
	inp := [][]int{{1, 2}, {4, 7}, {3, 6}}
	for i := 0; i < b.N; i++ {
		// Здесь выполняем функцию для тестирования
		fast(inp)
	}
}

func BenchmarkSimplestNTimesSlow(b *testing.B) {
	//inp := [][]int{{1, 2}, {4, 7}, {3, 6}}
	inp := [][]int{{4, 7}, {2, 4}, {9, 8}, {2, 7}, {4, 6}, {7, 3}, {9, 14}, {14, 9}, {5, 2}}
	for i := 0; i < b.N; i++ {
		// Здесь выполняем функцию для тестирования
		slow(inp)
	}
}

func Test_fast(t *testing.T) {
	type args struct {
		inpAll [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "correct small",
			args: args{
				[][]int{{1, 2}, {4, 7}, {3, 6}},
			},
			want: []int{1, 3, 0},
		},
		{
			name: "correct large",
			args: args{
				[][]int{{4, 7}, {2, 4}, {9, 8}, {2, 7}, {4, 6}, {7, 3}, {9, 14}, {14, 9}, {5, 2}},
			},
			want: []int{3, 2, 0, 2, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fast(tt.args.inpAll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slow(t *testing.T) {
	type args struct {
		inpAll [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "correct small",
			args: args{
				[][]int{{1, 2}, {4, 7}, {3, 6}},
			},
			want: []int{1, 3, 0},
		},
		{
			name: "correct large",
			args: args{
				[][]int{{4, 7}, {2, 4}, {9, 8}, {2, 7}, {4, 6}, {7, 3}, {9, 14}, {14, 9}, {5, 2}},
			},
			want: []int{3, 2, 0, 2, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slow(tt.args.inpAll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("slow() = %v, want %v", got, tt.want)
			}
		})
	}
}
