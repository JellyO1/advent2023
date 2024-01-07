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

	default:
		switch *part {
		case 2:
			run01Part2()
		default:
			run01()
		}
	}
}
