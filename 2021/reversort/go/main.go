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

	// Skip number of tests
	scanner.Scan()

	var cases int
	for scanner.Scan() {
		cases++

		parts := strings.Split(scanner.Text(), " ")

		intN, _ := strconv.Atoi(parts[0])
		intCost, _ := strconv.Atoi(parts[1])

		n := uint8(intN)
		cost := uint8(intCost)

		minCost := n - 1
		maxCost := n * (n + 1) / 2 - 1

		if cost < minCost || cost > maxCost {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", cases)
			continue
		}

		ints := buildList(cost, n)

		fmt.Printf("Case #%d: %v\n", cases, ints)
	}
}

func buildList(cost uint8, n uint8) []uint8 {
	var ints []uint8
	size := n

	var i uint8

	fmt.Println(size)
	
	for i = 1; i <= n; i++ {
		maxCost := size*(size+1)/2 - 1

		fmt.Println(cost, maxCost)

		if cost <= maxCost {
			ints = append(ints, i)
			fmt.Println(ints)
			cost--
			size--
		} else {
			ints = append(ints, buildWorstList(i, size)...)
			break
		}
	}

	return ints
}

func buildWorstList(last uint8, size uint8) []uint8 {
	worst := make([]uint8, size)

	elem := last + 1

	var i uint8

	for i = 0; i < size; i++ {
		worst[i] = elem
		elem += 2
	}

	return worst
}
