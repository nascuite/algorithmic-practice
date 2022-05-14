package main

import "fmt"

func main() {
	//res := twoSum([]int{3, 2, 4}, 6)
	//fmt.Println(res)
	fmt.Println(strStr("abc", "c"))
	fmt.Println(strStr("hello", "ll"))
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}
	return result
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	str := makeSliceFromStr(haystack)
	substr := makeSliceFromStr(needle)
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(substr); j++ {
			if j == len(substr)-1 && substr[j] == str[i+j] {
				return i
			}
			if i+j == len(str)-1 {
				return -1
			}
			if str[i+j] != substr[j] {
				break
			}
		}
	}
	return -1
}

func makeSliceFromStr(input string) []rune {
	var result []rune
	for _, s := range input {
		result = append(result, s)
	}
	return result
}
