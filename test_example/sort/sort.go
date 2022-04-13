package main

import (
	"fmt"
	"strconv"
)

func inpToSlice(inp string) []int {
	var result []int
	for _, s := range inp {
		i, _ := strconv.Atoi(string(s))
		result = append(result, i)
	}
	return result
}

func BubbleSort(ar []int) []int {
	for i := 0; i < len(ar); i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
	return ar
}

func main() {
	var inp string
	fmt.Scan(&inp)

	mapStr := make(map[int][]string)
	for i, val := range inp {
		strArr, exists := mapStr[i-1]
		var newStrArr []string
		if !exists {
			newStrArr = append(strArr, string(val))
		}

		for _, str := range strArr {
			for k, _ := range str {
				newStr := str[:k] + string(val) + str[k:]
				newStrArr = append(newStrArr, newStr)
			}
			newStrArr = append(newStrArr, str+string(val))
		}

		mapStr[i] = newStrArr
	}
	rez := mapStr[len(inp)-1]

	var inpSl []int
	for _, i := range rez {
		num, _ := strconv.Atoi(i)
		inpSl = append(inpSl, num)
	}

	inpSl = BubbleSort(inpSl)
	for _, s := range inpSl {
		fmt.Println(s)
	}
}
