package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type gameData struct {
	red   uint
	green uint
	blue  uint
}

type game struct {
	id   uint
	sets []gameData
}

func inputFromFile02(path string) ([]game, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	result := make([]game, 0)
	scanner := bufio.NewScanner(f)

	colorRegex := regexp.MustCompile("(\\d+) (red|green|blue)")

	for scanner.Scan() {
		line := scanner.Text()
		gameSplit := strings.Split(line, ":")
		gameId, err := strconv.ParseUint(strings.TrimLeft(gameSplit[0], "Game "), 10, 0)
		if err != nil {
			fmt.Printf("Failed to parse %s.\n", gameSplit[0])
			return nil, err
		}

		setSplit := strings.Split(gameSplit[1], ";")

		sets := make([]gameData, 0)
		for _, setstr := range setSplit {
			set := gameData{}
			colors := strings.Split(setstr, ",")
			for _, color := range colors {
				r := colorRegex.FindStringSubmatch(color)
				size, err := strconv.ParseUint(r[1], 10, 0)

				if err != nil {
					fmt.Printf("Failed to parse %s.\n", setstr)
					return nil, err
				}

				switch r[2] {
				case "blue":
					set.blue = uint(size)
				case "green":
					set.green = uint(size)
				case "red":
					set.red = uint(size)
				}
			}

			sets = append(sets, set)
		}

		g := game{
			id:   uint(gameId),
			sets: sets,
		}

		result = append(result, g)
	}

	return result, nil
}

func run02() {
	games, err := inputFromFile02("input/02.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}

	//var (
	//	totalRed   = uint(12)
	//	totalGreen = uint(13)
	//	totalBlue  = uint(14)
	//)

	result := uint(0)
	for _, g := range games {
		//possible := true
		var (
			minimumRed   = uint(0)
			minimumGreen = uint(0)
			minimumBlue  = uint(0)
		)

		for _, set := range g.sets {
			//if set.red > totalRed || set.blue > totalBlue || set.green > totalGreen {
			//	possible = false
			//	break
			//}
			if set.red > minimumRed {
				minimumRed = set.red
			}
			if set.blue > minimumBlue {
				minimumBlue = set.blue
			}
			if set.green > minimumGreen {
				minimumGreen = set.green
			}
		}

		//if possible {
		//	result += g.id
		//}
		result += minimumRed * minimumGreen * minimumBlue
	}

	fmt.Printf("%d\n", result)
}
