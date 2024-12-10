package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"slices"
	"strconv"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {

	var beforeMaps map[uint][]uint
	var afterMaps map[uint][]uint
	beforeMaps = make(map[uint][]uint)
	afterMaps = make(map[uint][]uint)

	var updates [][]uint

	for i, raw := range strings.Split(input, "\n\n") {
		if i == 0 {
			parseRules(&beforeMaps, &afterMaps, raw)
		}
		if i == 1 {
			parseUpdates(&updates, raw)
		}
	}

	var middles []uint

	if part2 {
		for _, update := range updates {
			valid, reordered := checkUpdateAndReorder(update, beforeMaps, afterMaps, true)
			if !valid {
				middleIndex := len(reordered) / 2
				middles = append(middles, reordered[middleIndex])
			}
		}

		return sumUint(middles)
	}

	for _, update := range updates {
		valid := checkUpdate(update, beforeMaps, afterMaps)
		if valid {
			middleIndex := len(update) / 2
			middles = append(middles, update[middleIndex])
		}
	}

	return sumUint(middles)
}

func checkUpdateAndReorder(update []uint, beforeMap map[uint][]uint, afterMap map[uint][]uint, valid bool) (bool, []uint) {
	var reordered []uint
	reordered = make([]uint, len(update))
	copy(reordered, update)

	for i, page := range update {
		beforeRules := beforeMap[page]
		afterRules := afterMap[page]

		if i < len(update) {
			beforeCheck, index := isValidOrderAccordingToRules(reordered[i+1:], beforeRules)
			if !beforeCheck {
				reordered[i], reordered[i+1+index] = reordered[i+1+index], reordered[i]
				valid = false
				return checkUpdateAndReorder(reordered, beforeMap, afterMap, valid)
			}
		}
		if i > 0 {
			afterCheck, index := isValidOrderAccordingToRules(reordered[:i], afterRules)
			if !afterCheck {
				reordered[i], reordered[i-1-index] = reordered[i-1-index], reordered[i]
				valid = false
				return checkUpdateAndReorder(reordered, beforeMap, afterMap, valid)
			}
		}
	}

	return valid, reordered
}

func checkUpdate(update []uint, beforeMap map[uint][]uint, afterMap map[uint][]uint) bool {
	valid := true

	for i, page := range update {
		beforeRules := beforeMap[page]
		afterRules := afterMap[page]

		if i < len(update) {
			beforeCheck, _ := isValidOrderAccordingToRules(update[i+1:], beforeRules)
			if !beforeCheck {
				valid = false
				break
			}
		}
		if i > 0 {
			afterCheck, _ := isValidOrderAccordingToRules(update[:i], afterRules)
			if !afterCheck {
				valid = false
				break
			}
		}
	}

	return valid
}

func isValidOrderAccordingToRules(pages []uint, rules []uint) (bool, int) {
	for i, page := range pages {
		if slices.Contains(rules, page) {
			return false, i
		}
	}
	return true, -1
}

func parseRules(beforeMap *map[uint][]uint, afterMap *map[uint][]uint, raw string) {
	for _, rule := range strings.Split(raw, "\n") {
		var before, after uint
		_, err := fmt.Sscanf(rule, "%d|%d", &before, &after)
		if err != nil {
			return
		}
		if !slices.Contains((*beforeMap)[after], before) {
			(*beforeMap)[after] = append((*beforeMap)[after], before)
		}
		if !slices.Contains((*afterMap)[before], after) {
			(*afterMap)[before] = append((*afterMap)[before], after)
		}
	}
}

func parseUpdates(updates *[][]uint, raw string) {
	for _, update := range strings.Split(raw, "\n") {
		strs := strings.Split(update, ",")
		pages := make([]uint, len(strs))
		for i := range strs {
			page, _ := strconv.Atoi(strs[i])
			pages[i] = uint(page)
		}
		*updates = append(*updates, pages)
	}
}

func sumUint(arr []uint) uint {
	sum := uint(0)
	for i := range arr {
		sum += arr[i]
	}
	return sum
}
