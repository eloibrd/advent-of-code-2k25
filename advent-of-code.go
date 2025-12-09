package main

import (
	day1 "eloibrd/advent-of-code-2k25/day-1"
	day2 "eloibrd/advent-of-code-2k25/day-2"
	day3 "eloibrd/advent-of-code-2k25/day-3"
	day4 "eloibrd/advent-of-code-2k25/day-4"
	day5 "eloibrd/advent-of-code-2k25/day-5"
	"flag"
	"fmt"
	"log"
	"log/slog"
)

func main() {
	day := flag.Int("day", 0, "Specify which day's puzzle to run (e.g. --day=1)")
	part := flag.Int("part", 0, "Specify which part to run (1 or 2) (e.g. --part=2)")
	flag.Parse()

	if *day == 0 || *part == 0 || *part > 2 {
		log.Fatalf("You must provide valid day and part")
		panic(1)
	}

	slog.Info(fmt.Sprintf("DAY %d | PART %d | Starting solving... ", *day, *part))

	var (
		result int
		err    error
	)

	switch *day {
	case 1:
		result, err = day1.SolvePassword(*part)
	case 2:
		result, err = day2.SolveGiftShop(*part)
	case 3:
		result, err = day3.SolveJoltage(*part)
	case 4:
		result, err = day4.SolveForklift(*part)
	case 5:
		result, err = day5.SolveCafeteriaInventory(*part)
	default:
		log.Fatalf("Day %d not yet implemented", *day)
	}
	if err != nil {
		panic(err)
	}

	slog.Info(fmt.Sprintf("DAY %d | PART %d | Solved : ", *day, *part), slog.Int("value", result))
}
