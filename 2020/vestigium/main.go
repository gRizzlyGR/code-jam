package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numTests, _ := strconv.Atoi(scanner.Text())

	for testCase := 0; testCase < numTests; testCase++ {
		scanner.Scan()

		n, _ := strconv.Atoi(scanner.Text())

		matrix := make([][]int, n)

		for r := 0; r < n; r++ {
			scanner.Scan()
			line := scanner.Text()
			arr := strings.Split(line, " ")

			ints := make([]int, n)
			for i, v := range arr {
				ints[i], _ = strconv.Atoi(v)
			}

			matrix[r] = ints
		}

		k, r, c := compute(matrix)
		fmt.Printf("Case #%d: %d %d %d\n", testCase+1, k, r, c)
	}
}

func compute(matrix [][]int) (int, int, int) {
	n := len(matrix)

	k := 0
	for i := range matrix {
		k += matrix[i][i]
	}

	r := 0
	for i := 0; i < n; i++ {
		dup := make(map[int]int)
		for j := 0; j < n; j++ {
			val := matrix[i][j]
			dup[val]++
		}

		if len(dup) != n {
			r++
		}
	}

	c := 0
	for j := 0; j < n; j++ {
		dup := make(map[int]int)
		for i := 0; i < n; i++ {
			val := matrix[i][j]
			dup[val]++
		}

		if len(dup) != n {
			c++
		}
	}

	return k, r, c
}
