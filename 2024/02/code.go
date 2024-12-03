package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"strconv"
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
	var reports [][]int

	var part int
	if part2 {
		part = 2
	} else {
		part = 1
	}
	fmt.Println("Day 2 - Part", part)

	var safeReportsCount = 0

	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)

		var levels []int
		for _, field := range fields {
			level, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error parsing", field)
			}
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}

	for _, report := range reports {
		safe := IsSafeReport(report)
		if safe {
			safeReportsCount++
		} else if part2 {
			for i := 0; i < len(report); i++ {
				if IsSafeReportWithDeletion(report, i) {
					safeReportsCount++
					break
				}
			}
		}
	}

	fmt.Println("Safe reports count", safeReportsCount)

	return safeReportsCount
}

func IsBadLevel(level int, nextLevel int, ascending bool) bool {
	distance := nextLevel - level
	//fmt.Println("Distance", distance)
	return (distance <= 0 && ascending) || (distance >= 0 && !ascending) || distance < -3 || distance > 3
}

func IsSafeReport(report []int) bool {
	var ascending = report[0] < report[1]
	var safe = true
	for i := 0; i < len(report)-1; i++ {
		if IsBadLevel(report[i], report[i+1], ascending) {
			safe = false
		}
	}

	return safe
}

func IsSafeReportWithDeletion(report []int, deleteIndex int) bool {
	copyReport := make([]int, len(report))
	copy(copyReport, report)

	if deleteIndex == len(copyReport)-1 {
		copyReport = copyReport[:deleteIndex]
	} else {
		copyReport = append(copyReport[:deleteIndex], copyReport[deleteIndex+1:]...)
	}

	return IsSafeReport(copyReport)
}
