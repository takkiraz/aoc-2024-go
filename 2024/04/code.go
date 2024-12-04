package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"regexp"
	"strings"
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
	xmasCount := 0

	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`XMAS`)
	reBackwards := regexp.MustCompile(`SAMX`)

	fmt.Println("Lines", len(lines))
	fmt.Println("Rows", len(lines[0]))
	fmt.Println("BorderX", len(lines[0])-len("XMAS"))
	fmt.Println("BorderY", len(lines)-len("XMAS"))

	for i, line := range lines {

		if !part2 {
			xmasCount += len(re.FindAllString(line, -1))
			xmasCount += len(reBackwards.FindAllString(line, -1))
		}

		for j, char := range line {
			if !part2 && char == 'X' {
				xmasCount += countXmas(lines, j, i)
			}
			if part2 && char == 'A' {
				xmasCount += countXmasPart2(lines, j, i)
			}
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return xmasCount
}

func countXmasPart2(lines []string, x, y int) int {
	count := 0

	minHeight := 1
	maxHeight := len(lines) - 2
	minWidth := 1
	maxWidth := len(lines[0]) - 2

	re := regexp.MustCompile("MAS|SAM")

	if x >= minWidth && x <= maxWidth && y >= minHeight && y <= maxHeight {
		diag1 := string(lines[y-1][x-1]) + string(lines[y][x]) + string(lines[y+1][x+1])
		diag2 := string(lines[y-1][x+1]) + string(lines[y][x]) + string(lines[y+1][x-1])

		if len(re.FindAllString(diag1, -1)) > 0 && len(re.FindAllString(diag2, -1)) > 0 {
			count++
		}

	}

	return count
}

func countXmas(lines []string, x, y int) int {
	borderY := len(lines) - len("XMAS")
	borderX := len(lines[0]) - len("XMAS")

	re := regexp.MustCompile("XMAS")

	count := 0

	substring := ""

	if x <= borderX {
		if y <= borderY {
			// unten rechts
			for i := 0; i < len("XMAS"); i++ {
				substring += string(lines[y+i][x+i])
			}

			count += len(re.FindAllString(substring, -1))
			substring = ""
		}

		if y >= len("XMAS")-1 {
			// oben rechts
			for i := 0; i < len("XMAS"); i++ {
				substring += string(lines[y-i][x+i])
			}

			count += len(re.FindAllString(substring, -1))
			substring = ""
		}
	}

	if x >= len("XMAS")-1 {
		if y <= borderY {
			// unten links
			for i := 0; i < len("XMAS"); i++ {
				substring += string(lines[y+i][x-i])
			}

			count += len(re.FindAllString(substring, -1))
			substring = ""
		}

		if y >= len("XMAS")-1 {
			// oben links
			for i := 0; i < len("XMAS"); i++ {
				substring += string(lines[y-i][x-i])
			}

			count += len(re.FindAllString(substring, -1))
			substring = ""
		}
	}

	// vertical
	if y <= borderY {
		for i := 0; i < len("XMAS"); i++ {
			substring += string(lines[y+i][x])
		}

		count += len(re.FindAllString(substring, -1))
		substring = ""
	}

	// vertical backwards
	if y >= len("XMAS")-1 {
		for i := 0; i < len("XMAS"); i++ {
			substring += string(lines[y-i][x])
		}

		count += len(re.FindAllString(substring, -1))
		substring = ""
	}

	return count
}
