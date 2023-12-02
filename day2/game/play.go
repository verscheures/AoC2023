package game

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Game struct {
	Round int
	Draws []map[string]int
}

type Output struct {
	sum   int
	power int
}

func Play(source string) {
	sum := 0
	power := 0
	readFile, err := os.Open(source)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for fileScanner.Scan() {

		game := parseLine(fileScanner.Text())
		out := validateGame(game, bag)
		sum += out.sum
		power += out.power
	}
	fmt.Println(sum)
	fmt.Println(power)
}

func parseLine(line string) Game {
	var game Game

	parts := strings.Split(line, ":")

	game.Round = getNumber(parts[0])
	draws := parseDraws(strings.Split(parts[1], ";"))
	game.Draws = draws
	return game
}

func parseDraws(sets []string) []map[string]int {
	var draws []map[string]int

	// split text
	for i := range sets {
		draw := make(map[string]int)
		counts := strings.Split(sets[i], ",")
		for j := range counts {
			num := getNumber(counts[j])
			color := getColor(counts[j])
			draw[color] = num

		}
		draws = append(draws, draw)
	}
	return draws
}

func getNumber(text string) int {
	re := regexp.MustCompile("[0-9]+")
	num, err := strconv.Atoi(re.FindString(text))
	check(err)
	return num
}

func getColor(text string) string {
	re := regexp.MustCompile("[a-z]+")
	return re.FindString(text)
}

func validateGame(game Game, bag map[string]int) Output {
	var output Output
	output.sum = 0
	output.power = 1
	powerbag := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	valid := true
	for i := range game.Draws {
		for k, v := range game.Draws[i] {
			if v > bag[k] {
				valid = false
			}
			if v > powerbag[k] {
				powerbag[k] = v
			}
		}
	}

	if valid {
		output.sum = game.Round
	}

	for _, v := range powerbag {
		output.power *= v
	}

	return output

}
