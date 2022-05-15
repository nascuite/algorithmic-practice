package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumberXOR([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	dict := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		if val, ok := dict[nums[i]]; !ok {
			dict[nums[i]] = val
		} else {
			delete(dict, nums[i])
		}
	}
	for k, _ := range dict {
		return k
	}
	return 0
}

func singleNumberXOR(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
