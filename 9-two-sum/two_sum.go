package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		find := target - nums[i]
		pos, exists := dict[find]
		if exists {
			result[0] = i
			result[1] = pos
			return result
		}
		dict[nums[i]] = i
	}
	return result
}
