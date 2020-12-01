package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input")

	if err != nil {
		panic(err)
	}

	input := toIntSlice(strings.Split(string(b), "\n"))

	for index, val := range input {
		for i := index + 1; i < len(input); i++ {
			if val+input[i] == 2020 {
				fmt.Println("answer pt1:", val*input[i])
			}
		}
	}

loop:
	for i := 0; i < len(input)/2; i++ {
		for j := len(input) - 1; j > len(input)/2; j-- {
			sum := input[i] + input[j]
			if sum < 2020 {
				for k := 0; k < len(input); k++ {
					if sum+input[k] == 2020 {
						fmt.Println("answer pt2:", input[i]*input[j]*input[k])
						break loop
					}
				}
			}
		}
	}

}

func toIntSlice(slice []string) []int {
	array := make([]int, len(slice))

	for index, s := range slice {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		array[index] = i
	}
	return array
}
