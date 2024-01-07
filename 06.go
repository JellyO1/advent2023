package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Duration       int
	RecordDistance int
}

func inputFromFile06(path string) ([]Race, error) {
	fs, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var races []Race
	scanner := bufio.NewScanner(fs)

	// Parse times
	scanner.Scan()
	timesLine := scanner.Text()

	times := make([]string, 0)
	for _, s := range strings.Split(strings.Split(timesLine, ":")[1], " ") {
		time := strings.TrimSpace(s)
		if time != "" {
			times = append(times, time)
		}
	}

	races = make([]Race, len(times))

	for i, time := range times {
		timeInt, err := strconv.Atoi(time)
		if err != nil {
			return nil, err
		}

		races[i] = Race{
			Duration: timeInt,
		}
	}

	scanner.Scan()
	distancesLine := scanner.Text()

	distances := make([]string, 0)
	for _, s := range strings.Split(strings.Split(distancesLine, ":")[1], " ") {
		distance := strings.TrimSpace(s)
		if distance != "" {
			distances = append(distances, distance)
		}
	}

	for i, distance := range distances {
		distanceInt, err := strconv.Atoi(distance)
		if err != nil {
			return nil, err
		}

		races[i].RecordDistance = distanceInt
	}

	fmt.Printf("%v\n", races)

	return races, nil
}

func calcRecordBeatingN(race Race) int {
	// Could use quadratic instead but the one below is simpler to visualize
	//duration := race.Duration
	//record := race.RecordDistance
	//
	//lowerBound := (float64(-duration) + math.Sqrt(float64(duration*duration-4*-1*-record))) / -2
	//upperBound := (float64(-duration) - math.Sqrt(float64(duration*duration-4*-1*-record))) / -2
	//
	//return int(upperBound - lowerBound + 1)

	i := 0

	for j := 0; j < race.Duration; j++ {
		if (race.Duration-j)*j > race.RecordDistance {
			i++
		}
	}

	fmt.Printf("race: %v, n: %d\n", race, i)
	return i
}

func run06() {
	races, err := inputFromFile06("input/06_test.txt")
	if err != nil {
		panic(err)
	}

	total := 1

	for _, race := range races {
		total *= calcRecordBeatingN(race)
	}

	fmt.Println(total)
}

func inputFromFile06Part2(path string) (*Race, error) {
	fs, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	race := Race{}
	scanner := bufio.NewScanner(fs)

	// Parse times
	scanner.Scan()
	timesLine := scanner.Text()

	time := ""
	for _, s := range strings.Split(strings.Split(timesLine, ":")[1], " ") {
		t := strings.TrimSpace(s)
		if t != "" {
			time += t
		}
	}

	timeInt, err := strconv.Atoi(time)
	if err != nil {
		return nil, err
	}
	race.Duration = timeInt

	scanner.Scan()
	distancesLine := scanner.Text()

	distance := ""
	for _, s := range strings.Split(strings.Split(distancesLine, ":")[1], " ") {
		d := strings.TrimSpace(s)
		if d != "" {
			distance += d
		}
	}

	distanceInt, err := strconv.Atoi(distance)
	if err != nil {
		return nil, err
	}
	race.RecordDistance = distanceInt

	fmt.Printf("%v\n", race)

	return &race, nil
}

func run06Part2() {
	race, err := inputFromFile06Part2("input/06.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(calcRecordBeatingN(*race))
}
