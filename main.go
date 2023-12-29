package main

import "flag"

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

	default:
		switch *part {
		case 2:
			run01Part2()
		default:
			run01()
		}
	}
}
