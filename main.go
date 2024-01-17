package main

import (
	"advent2023/day7"
	"advent2023/day7_p2"
	"flag"
	"fmt"
)

func main() {
	advent := flag.Uint("advent", 1, "The advent of code number to run")
	part := flag.Uint("part", 1, "The part to run")
	flag.Parse()

	switch *advent {
	case 2:
		run02()
	case 3:
		switch *part {
		case 2:
			run03Part2()
		default:
			run03()
		}
	case 4:
		switch *part {
		case 2:
			run04Part2()
		default:
			run04()
		}
	case 5:
		switch *part {
		case 2:
			run05Part2()
		default:
			run05()
		}
	case 6:
		switch *part {
		case 2:
			run06Part2()
		default:
			run06()
		}
	case 7:
		switch *part {
		case 2:
			fmt.Print(day7_p2.Run("input/07.txt"))
		default:
			fmt.Print(day7.Run("input/07.txt"))
		}

	default:
		switch *part {
		case 2:
			run01Part2()
		default:
			run01()
		}
	}
}
