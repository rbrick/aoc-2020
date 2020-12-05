package main

import (
	"bufio"
	"fmt"
	"os"
)

func Parse(line string) (int, int) {
	minRow, maxRow := 0, 127
	minColumn, maxColumn := 0, 7

	row := 0
	column := 0

	for _, c := range line {
		// so for here you are basically just setting the median to min/max
		if c == 'F' {
			maxRow = (maxRow-minRow)/2 + minRow
		} else if c == 'B' {
			minRow = minRow + (maxRow-minRow)/2 + 1
		} else if c == 'L' {
			maxColumn = (maxColumn-minColumn)/2 + minColumn
		} else if c == 'R' {
			minColumn = minColumn + (maxColumn-minColumn)/2 + 1
		}
	}

	row = maxRow
	column = maxColumn

	return row, column
}

func main() {
	f, err := os.Open("input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	highestSeat := 0

	var seats []int

	for scanner.Scan() {
		row, column := Parse(scanner.Text())
		seatId := row*8 + column

		if seatId > highestSeat {
			highestSeat = seatId
		}

		seats = append(seats, seatId)
	}

	seatId := 0

	for _, seat := range seats {
		// search for a seat that doesn't exist AKA our seat
		if ok, idx := contains(seats, seat-1); !ok {
			// if we find a seat that doesn't exist,
			// check if the our seat minus one exist (so current current existing seat -> our seat -> existing seat)
			// if we find a pattern like this, it is our seat
			if ok, idx = contains(seats, seat-2); ok {
				seatId = seats[idx] + 1
			}
		}
	}

	fmt.Println("part 1:", highestSeat)
	fmt.Println("part 2:", seatId)
}

func contains(a []int, e int) (bool, int) {
	for idx, v := range a {
		if v == e {
			return true, idx
		}
	}
	return false, -1
}
