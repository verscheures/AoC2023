package parse

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	sum := 0
	readFile, err := os.Open(source)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// fileLines = append(fileLines, fileScanner.Text())
		// sum += getNumber(fileScanner.Text())
		sum += getAllNumbersCorrect(fileScanner.Text())
		fmt.Println(sum)
	}
	fmt.Println(sum)
}

func getNumber(line string) int {
	numbers := "1234567890"
	fmt.Println(line)
	firstLoc := strings.IndexAny(line, numbers)
	lastLoc := strings.LastIndexAny(line, numbers)
	first, err := strconv.Atoi(string([]rune(line)[firstLoc]))
	check(err)
	last, err := strconv.Atoi(string([]rune(line)[lastLoc]))
	check(err)
	fmt.Println(first, " ", last)
	return first*10 + last
}

func getAllNumbers(line string) int {
	re := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")

	matches := re.FindAllString(line, -1)
	fmt.Print(line, " ")
	fmt.Print(hit2Num(matches[0]), " = ", matches[0])
	fmt.Println(hit2Num(matches[len(matches)-1]), " = ", matches[len(matches)-1])
	return hit2Num(matches[0])*10 + hit2Num(matches[len(matches)-1])
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
		// fmt.Println(i, " ", k, v)
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
	if first.value == -1 || last.value == -1 {
		fmt.Println("stinker", line)
	}
	fmt.Println(line, " ", first.value*10+last.value)
	return first.value*10 + last.value
}

func hit2Num(hit string) int {
	number, err := strconv.Atoi(hit)
	if err != nil {
		switch hit {
		case "one":
			return 1
		case "two":
			return 2
		case "three":
			return 3
		case "four":
			return 4
		case "five":
			return 5
		case "six":
			return 6
		case "seven":
			return 7
		case "eight":
			return 8
		case "nine":
			return 9
		}
	}
	return number
}
