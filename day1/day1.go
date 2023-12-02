package day1

import (
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

func Day1() {
	var total = 0
	dat, err := os.ReadFile("day1/puzzleInput")
	check(err)
	fullFile := string(dat)
	splitedFile := strings.Split(fullFile, "\n")
	// iterate lines
	for _, line := range splitedFile {
		var digitInLine []string
		// iterate chars
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			match, _ := regexp.MatchString(`\d`, string(char))
			if match {
				check(err)
				digitInLine = append(digitInLine, char)
			}
		}
		// create number
		numberInLine := digitInLine[0] + digitInLine[len(digitInLine)-1]
		number, err := strconv.Atoi(numberInLine)
		check(err)
		total = total + number
	}
	fmt.Println("Total : ", total)
}
