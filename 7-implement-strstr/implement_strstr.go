package main

import "fmt"

func main() {
	fmt.Println(strStr("abc", "c"))
	fmt.Println(strStr("hello", "ll"))
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
