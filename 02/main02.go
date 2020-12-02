package main

import (
	"bufio"
	"os"
)

func main() {
	f, err := os.Open("input")

	if err != nil {
		panic(err)
	}

	_ = bufio.NewScanner(f)
}
