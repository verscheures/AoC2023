package parse

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type hit struct {
	index int
	value int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Parse(source string) {
	part1 := 0
	part2 := 0

	readFile, err := os.Open(source)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}
		part1 += getNumber(line)
		part2 += getAllNumbersCorrect(line)
	}
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func getNumber(line string) int {
	numbers := "1234567890"
	firstLoc := strings.IndexAny(line, numbers)
	lastLoc := strings.LastIndexAny(line, numbers)
	first, err := strconv.Atoi(string([]rune(line)[firstLoc]))
	check(err)
	last, err := strconv.Atoi(string([]rune(line)[lastLoc]))
	check(err)
	return first*10 + last
}

func getAllNumbersCorrect(line string) int {
	numbers := map[string]int{
		"1":     1,
		"one":   1,
		"2":     2,
		"two":   2,
		"3":     3,
		"three": 3,
		"4":     4,
		"four":  4,
		"5":     5,
		"five":  5,
		"6":     6,
		"six":   6,
		"7":     7,
		"seven": 7,
		"8":     8,
		"eight": 8,
		"9":     9,
		"nine":  9,
	}
	first := hit{len(line), 0}
	last := hit{-1, 0}
	for k, v := range numbers {
		f := strings.Index(line, k)
		l := strings.LastIndex(line, k)
		if f == -1 {
			continue
		}
		if f < first.index {
			first.index = f
			first.value = v
		}
		if l > last.index {
			last.index = l
			last.value = v
		}
	}
	return first.value*10 + last.value
}
