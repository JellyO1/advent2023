package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Map struct {
	Name        string
	Destination int
	Source      int
	Range       int
}

func (m Map) Convert(source int) int {
	if source < m.Source || source > m.Source+(m.Range-1) {
		return source
	}

	diff := m.Destination - m.Source
	return source + diff
}

func (m Map) String() string {
	return fmt.Sprintf("%s\n d:%d, s:%d, r:%d\n", m.Name, m.Destination, m.Source, m.Range)
}

func inputFromFile05(path string) ([]int, [][]Map, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	var seeds []int

	mapOfMaps := make([][]Map, 7)
	i := 0

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(text, "seeds:") {
			numStrings := strings.Split(text, " ")[1:]

			seeds = make([]int, len(numStrings))
			for i, ns := range numStrings {
				n, err := strconv.Atoi(ns)
				if err != nil {
					return nil, nil, err
				}
				seeds[i] = n
			}
			continue
		}

		if strings.Contains(text, "map") {
			name := text
			maps := make([]Map, 0)

			for scanner.Scan() {
				m := Map{
					Name: name,
				}

				text = scanner.Text()
				if text == "" {
					break
				}

				numStrings := strings.Split(text, " ")
				for j, numString := range numStrings {
					n, err := strconv.Atoi(numString)
					if err != nil {
						return nil, nil, err
					}

					switch j {
					case 0:
						m.Destination = n
					case 1:
						m.Source = n
					case 2:
						m.Range = n
					}
				}

				maps = append(maps, m)
			}

			mapOfMaps[i] = maps
			i++
		}
	}

	return seeds, mapOfMaps, nil
}

func run05() {
	seeds, mapOfMaps, err := inputFromFile05("input/05.txt")
	if err != nil {
		panic(err)
	}

	for _, mm := range mapOfMaps {
		for i, seed := range seeds {
			for _, m := range mm {
				conv := m.Convert(seed)
				if conv != seed {
					seeds[i] = conv
					break
				}
			}
		}
		fmt.Printf("%s, seeds: %v\n", mm[0].Name, seeds)
	}

	lowest := math.MaxInt
	for _, s := range seeds {
		if s < lowest {
			lowest = s
		}
	}

	fmt.Printf("lowest: %d\n", lowest)
}

type SeedRange struct {
	Start int
	Range int
}

func run05Part2() {
	seedsAndRanges, mapOfMaps, err := inputFromFile05("input/05.txt")
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	c := make(chan int)
	for i := 0; i < len(seedsAndRanges); i += 2 {
		sr := SeedRange{
			Start: seedsAndRanges[i],
			Range: seedsAndRanges[i+1],
		}

		wg.Add(1)
		go computeLowestForRange(mapOfMaps, sr, c, wg)
	}

	fmt.Printf("number of goroutines: %d\n", len(seedsAndRanges)/2)

	go func() {
		wg.Wait()
		close(c)
	}()

	//fmt.Printf("lowest: %d\n", seeds)

	lowest := math.MaxInt
	for s := range c {
		if s < lowest {
			lowest = s
		}
	}

	fmt.Printf("lowest: %d\n", lowest)
}

func computeLowestForRange(mapOfMaps [][]Map, seedRange SeedRange, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	seeds := make([]int, seedRange.Range)
	for _, mm := range mapOfMaps {
		for i := seedRange.Start; i < seedRange.Start+seedRange.Range; i++ {
			var seed int
			if seeds[i-seedRange.Start] == 0 {
				seed = i
			} else {
				seed = seeds[i-seedRange.Start]
			}
			for _, m := range mm {
				conv := m.Convert(seed)
				if conv != seed {
					seeds[i-seedRange.Start] = conv
					break
				}
			}
		}
	}

	lowest := math.MaxInt
	for _, s := range seeds {
		if s < lowest {
			lowest = s
		}
	}

	fmt.Printf("range: %v, lowest: %d\n", seedRange, lowest)
	c <- lowest
}
