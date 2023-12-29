package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winningNumbers []int
	numbers        []int
}

func inputFromFile04(path string) ([]Card, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	result := make([]Card, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		s := strings.Split(t, ":")
		s = strings.Split(s[1], "|")

		card := Card{}

		// parse winning numbers
		numbersStrings := strings.Split(strings.TrimSpace(s[0]), " ")
		i := 0
		for _, numbersString := range numbersStrings {
			if numbersString != "" {
				numbersStrings[i] = numbersString
				i++
			}
		}
		numbersStrings = numbersStrings[:i]

		card.winningNumbers = make([]int, len(numbersStrings))
		for i, numberString := range numbersStrings {
			n, err := strconv.Atoi(numberString)
			if err != nil {
				return nil, err
			}

			card.winningNumbers[i] = n
		}

		// parse numbers
		numbersStrings = strings.Split(strings.TrimSpace(s[1]), " ")

		i = 0
		for _, numbersString := range numbersStrings {
			if numbersString != "" {
				numbersStrings[i] = numbersString
				i++
			}
		}
		numbersStrings = numbersStrings[:i]

		card.numbers = make([]int, len(numbersStrings))
		for i, numberString := range numbersStrings {
			n, err := strconv.Atoi(numberString)
			if err != nil {
				return nil, err
			}

			card.numbers[i] = n
		}

		result = append(result, card)
	}

	return result, nil
}

func run04() {
	cards, err := inputFromFile04("input/04.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	for i, card := range cards {
		// The first card is only worth 1 point, so 2^0
		winningNumberTimes := -1
		for _, winning := range card.winningNumbers {
			for _, number := range card.numbers {
				if winning == number {
					winningNumberTimes++
				}
			}
		}

		if winningNumberTimes < 0 {
			continue
		}

		fmt.Printf("card %d: winning:%d, value:%.0f\n", i+1, winningNumberTimes+1, math.Pow(2, float64(winningNumberTimes)))
		total += int(math.Pow(2, float64(winningNumberTimes)))
	}

	fmt.Printf("total: %d\n", total)
}

func run04Part2() {
	cards, err := inputFromFile04("input/04.txt")
	if err != nil {
		panic(err)
	}

	copies := make(map[int]int, len(cards))

	for i, card := range cards {
		winningNumberTimes := 0
		for _, winning := range card.winningNumbers {
			for _, number := range card.numbers {
				if winning == number {
					copies[i+winningNumberTimes+1] += copies[i] + 1
					winningNumberTimes++
				}
			}
		}

	}

	fmt.Printf("%v\n", copies)

	total := len(cards)

	for _, v := range copies {
		total += v
	}

	fmt.Printf("total: %d", total)
}
