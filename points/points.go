package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var cnt, inputN, inputT int
	var inp string
	fmt.Scan(&cnt)

	inpAllNums := make([][]int, cnt+1)

	for i := 0; i <= cnt; i++ {
		inpNums := make([]int, 2)
		fmt.Scan(&inp)
		inputN, _ = strconv.Atoi(inp)
		inpNums[0] = inputN

		fmt.Scan(&inp)
		inputT, _ = strconv.Atoi(inp)
		inpNums[1] = inputT

		inpAllNums[i] = inpNums
	}

	//results := fast(inpAllNums)
	results := slow(inpAllNums)
	for _, result := range results {
		println(result)
	}

}

func fast(inpAll [][]int) []int {
	res := make([]int, len(inpAll))

	for i := 0; i < len(inpAll); i++ {
		inpLine := inpAll[i]
		if inpLine[0] >= inpLine[1] {
			res[i] = 0
		} else {
			res[i] = factorialMod(inpLine[0], float64(inpLine[1]))
		}
	}

	return res
}

func slow(inp [][]int) []int {
	res := make([]int, len(inp))

	for i := 0; i < len(inp); i++ {
		inpLine := inp[i]
		f := factorial(inpLine[0])
		res[i] = int(math.Mod(float64(f), float64(inpLine[1])))
	}

	return res
}

func factorialMod(n int, divider float64) int {
	var res float64 = 1

	for i := 2; i <= n; i++ {
		res = math.Mod(res*float64(i), divider)
	}

	return int(res)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
