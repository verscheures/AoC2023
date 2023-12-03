package fix

import (
	"bufio"

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

type Engine struct {
	Lines [3]string
	Sum   int
	Ratio int
}

func (e *Engine) Init(source string) {
	readFile, err := os.Open(source)

	check(err)
	defer readFile.Close()
	FileScanner := bufio.NewScanner(readFile)
	FileScanner.Split(bufio.ScanLines)

	FileScanner.Scan()

	e.Lines[1] = FileScanner.Text()
	e.Lines[0] = strings.Repeat(".", len(e.Lines[1]))
	e.Sum = 0

	for FileScanner.Scan() {
		e.Lines[2] = FileScanner.Text()
		sum, ratio := e.scanLine()
		e.Sum += sum
		e.Ratio += ratio
		//fmt.Println("Sum:", e.Sum)
		e.Lines[0] = e.Lines[1]
		e.Lines[1] = e.Lines[2]
	}
	e.Lines[2] = strings.Repeat(".", len(e.Lines[1]))
	sum, ratio := e.scanLine()
	e.Sum += sum
	e.Ratio += ratio

}

func (e *Engine) scanLine() (int, int) {
	count := 0
	ratio := 0
	reSymbol := regexp.MustCompile(`[^0-9\.]`)
	reNumber := regexp.MustCompile(`[0-9]+`)
	hits := reSymbol.FindAllIndex([]byte(e.Lines[1]), -1)

	for _, h := range hits {
		var startHits []int

		//fmt.Println("Starhits", startHits)
		for i := 0; i < 3; i++ {
			numbers := reNumber.FindAllIndex([]byte(e.Lines[i]), -1)
			for _, n := range numbers {
				// checking part1
				if n[0] <= h[1] && n[1] >= h[0] {

					number, err := strconv.Atoi(e.Lines[i][n[0]:n[1]])
					check(err)
					count += number

					if e.Lines[1][h[0]:h[1]] == "*" {
						startHits = append(startHits, number)

					}
				}
			}

		}
		if len(startHits) == 2 {
			ratio += startHits[0] * startHits[1]
		}
	}
	return count, ratio
}
