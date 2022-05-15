package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
}

func plusOne(digits []int) []int {
	for i := len(digits); i >= 0; i-- {
		if i == 0 {
			result := make([]int, len(digits)+1)
			result[0] = 1
			copy(result[1:], digits)
			return result
		}
		if digits[i-1] < 9 {
			digits[i-1] = digits[i-1] + 1
			break
		} else {
			digits[i-1] = 0
			continue
		}
	}
	return digits
}
