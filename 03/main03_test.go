package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

var testInput []string

func init() {
	b, err := ioutil.ReadFile("input")

	if err != nil {
		panic(err)
	}

	testInput = strings.Split(string(b), "\n")
}

func TestPartOne(t *testing.T) {
	expectedResult := 7

	result := traverse(testInput, 3, 1)

	if result != expectedResult {
		t.Errorf("failed test, expected result => %d, got %d", expectedResult, result)
	}
}

func TestPartTwo(t *testing.T) {
	expectedResult := 336

	checkedSlopes := []int{
		traverse(testInput, 3, 1),
		traverse(testInput, 1, 1),
		traverse(testInput, 5, 1),
		traverse(testInput, 7, 1),
		traverse(testInput, 1, 2),
	}

	result := 1

	for _, v := range checkedSlopes {
		result = result * v
	}

	if result != expectedResult {
		t.Errorf("failed test, expected result => %d, got %d", expectedResult, result)
	}
}
