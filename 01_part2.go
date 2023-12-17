package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func run01Part2() {
	input, err := inputFromFile("input/01.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var firstDigit, lastDigit int
	var sum int

	exp, err := regexp.Compile("(\\d|one|two|three|four|five|six|seven|eight|nine).*?(?:.*(\\d|one|two|three|four|five|six|seven|eight|nine))?")
	for _, s := range input {
		matches := exp.FindStringSubmatch(s)
		fmt.Printf("%q\n", matches)
		for i := 1; i < len(matches); i++ {
			match := matches[i]

			val, err := strconv.Atoi(match)
			if err != nil {
				switch match {
				case "one":
					val = 1
				case "two":
					val = 2
				case "three":
					val = 3
				case "four":
					val = 4
				case "five":
					val = 5
				case "six":
					val = 6
				case "seven":
					val = 7
				case "eight":
					val = 8
				case "nine":
					val = 9
				default:
					continue
				}
			}

			if firstDigit == 0 {
				firstDigit = val
				lastDigit = val
			} else {
				lastDigit = val
			}
		}

		//fmt.Printf("%d%d\n", firstDigit, lastDigit)
		val, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))

		if err != nil {
			panic(err)
		}

		sum += val
		firstDigit, lastDigit = 0, 0
	}

	fmt.Printf("%d", sum)
}
