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
	n, _ := strconv.Atoi(scanner.Text())

	for testCase := 0; testCase < n; testCase++ {
		scanner.Scan()
		line := scanner.Text()
		elems := strings.Split(line, "")

		ints := make([]int, len(elems))
		for e := range elems {
			ints[e], _ = strconv.Atoi(elems[e])
		}

		str := ""

		for _, d := range ints {
			for k := 0; k < d; k++ {
				str += "("
			}

			str += strconv.Itoa(d)

			for k := 0; k < d; k++ {
				str += ")"
			}
		}

		str = reduce(str)

		fmt.Printf("Case #%d: %s\n", testCase+1, str)
	}
}

func reduce(str string) string {
	for {
		newStr := strings.Replace(str, ")(", "", -1)
		if newStr == str {
			return str
		}
		str = newStr
	}
}