package main

import (
	"bufio"
	"fmt"
	"os"
)

func Count(line string, group map[rune]int) {
	for _, c := range line {
		group[c]++
	}
}

func main() {
	f, err := os.Open("input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	group := map[rune]int{}
	memberCount := 0
	var groups []map[rune]int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			group['M'] = memberCount
			memberCount = 0

			groups = append(groups, group)
			group = map[rune]int{}
		} else {
			memberCount++
		}
		Count(line, group)
	}

	group['M'] = memberCount
	groups = append(groups, group)

	totalAnswers := 0
	allYes := 0
	for _, group := range groups {
		totalAnswers += len(group) - 1
		memberCount := group['M']

		for r, v := range group {
			if r != 'M' && v == memberCount {
				allYes++
			}
		}
	}

	fmt.Println("part1:", totalAnswers)
	fmt.Println("part2:", allYes)
}
