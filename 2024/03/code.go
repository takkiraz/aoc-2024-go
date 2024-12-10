package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"regexp"
)

func main() {
	aoc.Harness(run)
}

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
