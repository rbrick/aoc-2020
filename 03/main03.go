package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func traverse(s []string, xSlope, ySlope int) int {
	x := xSlope
	count := 0
	for i := ySlope; i < len(s); i += ySlope {
		line := s[i]
		lineIndex := x % len(line)

		if line[lineIndex] == '#' {
			count++
		}

		x += xSlope
	}
	return count
}

func main() {
	b, err := ioutil.ReadFile("input")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(b), "\n")

	checkedSlopes := []int{
		traverse(input, 1, 1),
		traverse(input, 3, 1),
		traverse(input, 5, 1),
		traverse(input, 7, 1),
		traverse(input, 1, 2),
	}

	fmt.Println("part 1 answer:", checkedSlopes[1])

	result := 1
	for _, v := range checkedSlopes {
		result = result * v
	}

	fmt.Println("part 2 answer:", result)
}
