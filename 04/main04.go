package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type FieldCheck func(string) bool

func Range(min, max int) FieldCheck {
	return func(s string) bool {
		i, _ := strconv.Atoi(s)
		return i >= min && i <= max
	}
}

func Pattern(pattern string) FieldCheck {
	return func(s string) bool {
		pattern := regexp.MustCompile(pattern)
		return pattern.MatchString(s)
	}
}

func StrLen(length int) FieldCheck {
	return func(s string) bool {
		return len(s) == length
	}
}

func StrEqualsAny(any ...string) FieldCheck {
	return func(s string) bool {
		equals := false
		for _, check := range any {
			if s == check {
				equals = true
			}
		}
		return equals
	}
}

func Height(s string) bool {
	pattern := regexp.MustCompile("([0-9]+)(in|cm)")
	if pattern.MatchString(s) {
		matches := pattern.FindAllStringSubmatch(s, 1)
		i, _ := strconv.Atoi(matches[0][1])

		if matches[0][2] == "cm" {
			return i >= 150 && i <= 193
		} else {
			return i >= 59 && i <= 76
		}

	}
	return false
}

func Compound(checks ...FieldCheck) FieldCheck {
	return func(s string) bool {
		for _, check := range checks {
			if !check(s) {
				return false
			}
		}
		return true
	}
}

var (
	requiredFields = map[string]FieldCheck{
		"byr": Range(1920, 2002),
		"iyr": Range(2010, 2020),
		"eyr": Range(2020, 2030),
		"hgt": Height,
		"hcl": Compound(Pattern("#[a-f0-9]{6}"), StrLen(7)),
		"ecl": StrEqualsAny("amb", "blu", "brn", "gry", "grn", "hzl", "oth"),
		"pid": Compound(Pattern("[0-9]{9}"), StrLen(9)),
	}
	pattern = regexp.MustCompile("([a-z]{3}):([\\w#]+)")
)

type PassportEntry map[string]string

func (entry PassportEntry) HasFields() bool {
	for field, _ := range requiredFields {
		if _, ok := entry[field]; !ok {
			return false
		}
	}
	return true
}

func (entry PassportEntry) Valid() bool {
	for field, check := range requiredFields {
		if _, ok := entry[field]; !ok {
			return false
		} else {
			if !check(entry[field]) {
				return false
			}
		}
	}
	return true
}

func ParseEntry(line string, entry PassportEntry) {
	matches := pattern.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		entry[match[1]] = match[2]
	}
}

func main() {

	f, err := os.Open("input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	currentEntry := PassportEntry{}

	var entries []PassportEntry

	fieldsPresent := 0
	validEntry := 0

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			if currentEntry.HasFields() {
				fieldsPresent++

				if currentEntry.Valid() {
					validEntry++
				}

			}

			entries = append(entries, currentEntry)
			currentEntry = PassportEntry{}
			continue
		}
		ParseEntry(txt, currentEntry)
	}

	fmt.Println("pt1 answer:", fieldsPresent)
	fmt.Println("pt2 answer:", validEntry)
}
