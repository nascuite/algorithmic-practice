package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {
	lp := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[lp] {
			lp += 1
			nums[lp] = nums[i]
		}
	}
	return lp + 1
}
