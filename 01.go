package main

import (
	"fmt"
	"os"
	"strconv"
)

func run01() {
	input, err := inputFromFile("input/01.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var firstDigit, lastDigit int
	var sum int
	for _, s := range input {
		//fmt.Printf("%s\n", s)
		for i := 0; i < len(s); i++ {
			r := s[i]

			val, err := strconv.Atoi(string(r))
			if err != nil {
				// Not a number ignore
				continue
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
