package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"regexp"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	if part2 {
		filter := regexp.MustCompile(`(?s)don't\(\).*?do\(\)|$`)

		input = filter.ReplaceAllString(input, "")
	}

	var result = 0

	for _, match := range re.FindAllStringSubmatch(input, -1) {
		var left, right int
		_, err := fmt.Sscanf(match[0], "mul(%d,%d)", &left, &right)
		if err != nil {
			return nil
		}

		result += left * right
	}

	return result
}
