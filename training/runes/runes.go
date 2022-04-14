package runes

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var out strings.Builder
	fmt.Scan(&input)
	for i, s := range input {
		if i > 0 && i < 4 {
			out.WriteRune(s)
		}
		if i > 4 {
			break
		}
	}
	fmt.Println(out.String())
}
