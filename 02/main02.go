package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile("([0-9]+)-([0-9]+)\\s([\\w]):\\s([\\w]+)")

type PasswordPolicy struct {
	Min, Max int
	Matching string
}

func (policy *PasswordPolicy) Validate(password string) bool {
	count := strings.Count(password, policy.Matching)
	return count >= policy.Min && count <= policy.Max
}

func (policy *PasswordPolicy) Validate2(password string) bool {
	score := 0

	if password[policy.Min-1] == policy.Matching[0] {
		score++
	}

	if password[policy.Max-1] == policy.Matching[0] {
		score++
	}
	return score == 1

}

type PasswordEntry struct {
	Policy   *PasswordPolicy
	Password string
}

func Read(line string) *PasswordEntry {
	matches := regex.FindAllStringSubmatch(line, -1)

	min, _ := strconv.Atoi(matches[0][1])
	max, _ := strconv.Atoi(matches[0][2])
	matching := matches[0][3]
	password := matches[0][4]

	return &PasswordEntry{
		Policy: &PasswordPolicy{
			Min:      min,
			Max:      max,
			Matching: matching,
		},
		Password: password,
	}
}

func main() {
	f, err := os.Open("input")

	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	validPasswords1 := 0
	validPasswords2 := 0

	for s.Scan() {
		entry := Read(s.Text())

		if entry.Policy.Validate(entry.Password) {
			validPasswords1++
		}

		if entry.Policy.Validate2(entry.Password) {
			validPasswords2++
		}
	}

	fmt.Println("answer pt1:", validPasswords1)
	fmt.Println("answer pt2:", validPasswords2)
}
