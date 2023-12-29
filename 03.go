package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Schematic struct {
	Width  int
	Height int
	Data   [][]rune
}

func inputFromFile03(path string) (*Schematic, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	schematic := Schematic{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("%s\n", text)
		schematic.Height++
		if schematic.Width == 0 {
			schematic.Width = len(scanner.Bytes())
		}
	}

	schematic.Data = make([][]rune, schematic.Height)

	if _, err = f.Seek(0, 0); err != nil {
		return nil, err
	}
	scanner = bufio.NewScanner(f)

	i := 0
	for scanner.Scan() {
		schematic.Data[i] = make([]rune, schematic.Width)
		text := scanner.Text()
		for j, b := range text {
			schematic.Data[i][j] = b
		}
		i++
	}

	return &schematic, nil
}

func run03() {
	schematic, _ := inputFromFile03("input/03.txt")

	total := 0

	for h := 0; h < schematic.Height; h++ {
		number := ""
		numStartIndex := -1
		numEndIndex := -1

		for w := 0; w < schematic.Width; w++ {
			r := schematic.Data[h][w]

			_, err := strconv.Atoi(string(r))
			if err == nil {
				if numStartIndex == -1 {
					numStartIndex = w - 1
				}

				number += string(r)
			}

			// we found a character that isn't a number or reached the end
			// and we are currently processing a number.
			if (err != nil || w == schematic.Width-1) && numStartIndex != -1 {
				numEndIndex = w

				// check for a symbol around the number
				for y := max(h-1, 0); y <= min(h+1, schematic.Height-1); y++ {
					foundSymbol := false
					for x := max(numStartIndex, 0); x <= min(numEndIndex, schematic.Width-1); x++ {
						r := schematic.Data[y][x]
						// Symbol
						if _, err = strconv.Atoi(string(r)); err != nil && r != '.' {
							n, err := strconv.Atoi(number)
							if err != nil {
								fmt.Fprint(os.Stderr, err)
								panic(err)
							}
							total += n
							foundSymbol = true
							break
						}
					}

					if foundSymbol {
						break
					}
				}

				numStartIndex = -1
				numEndIndex = -1
				number = ""
			}
		}
	}

	fmt.Printf("total: %d\n", total)
}

func run03Part2() {
	schematic, _ := inputFromFile03("input/03.txt")

	gearNumbers := make(map[string][]int)

	for h := 0; h < schematic.Height; h++ {
		number := ""
		numStartIndex := -1
		numEndIndex := -1

		for w := 0; w < schematic.Width; w++ {
			r := schematic.Data[h][w]

			_, err := strconv.Atoi(string(r))
			if err == nil {
				if numStartIndex == -1 {
					numStartIndex = w - 1
				}

				number += string(r)
			}

			if (err != nil || w == schematic.Width-1) && numStartIndex != -1 {
				numEndIndex = w

				// check for a symbol
				for y := max(h-1, 0); y <= min(h+1, schematic.Height-1); y++ {
					for x := max(numStartIndex, 0); x <= min(numEndIndex, schematic.Width-1); x++ {
						r := schematic.Data[y][x]
						// Symbol
						if _, err = strconv.Atoi(string(r)); err != nil && r == '*' {
							n, err := strconv.Atoi(number)
							if err != nil {
								fmt.Fprint(os.Stderr, err)
								panic(err)
							}

							symbolKey := fmt.Sprintf("x%dy%d", x, y)

							if gearNumbers[symbolKey] == nil {
								gearNumbers[symbolKey] = make([]int, 0)
							}

							gearNumbers[symbolKey] = append(gearNumbers[symbolKey], n)
						}
					}
				}

				numStartIndex = -1
				numEndIndex = -1
				number = ""
			}
		}
	}

	var total int
	for _, ints := range gearNumbers {
		if len(ints) != 2 {
			continue
		}

		total += ints[0] * ints[1]
	}

	fmt.Printf("total: %d", total)
}
