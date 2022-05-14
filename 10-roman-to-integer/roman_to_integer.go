package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	dict := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var value int
	var oldSymb rune
	for _, symb := range s {
		if (symb == 'X' || symb == 'V') && oldSymb == 'I' {
			value += -1*dict[oldSymb] + dict[symb] - dict[oldSymb]
		} else if (symb == 'L' || symb == 'C') && oldSymb == 'X' {
			value += -1*dict[oldSymb] + dict[symb] - dict[oldSymb]
		} else if (symb == 'D' || symb == 'M') && oldSymb == 'C' {
			value += -1*dict[oldSymb] + dict[symb] - dict[oldSymb]
		} else {
			value += dict[symb]
		}
		oldSymb = symb
	}
	return value
}
