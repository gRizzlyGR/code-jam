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
		parts := strings.Split(scanner.Text(), " ")
		
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		mural := parts[2]
		
		var cost int
		var prev rune
		
		for _, letter := range mural {
			switch letter {
			case 'C':
				{

					if prev == 'J' {
						cost += y
					}
					prev = letter
				}
			case 'J':
				{

					if prev == 'C' {
						cost += x
					}
					prev = letter
				}
			default:
				{
					// Ignore '?'
				}
			}
		}
		cases++
		fmt.Printf("Case #%d: %d\n", cases, cost)
	}
}
