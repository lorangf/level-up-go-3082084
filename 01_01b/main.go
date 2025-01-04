package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	bday, err := time.Parse(expectedFormat, target)
	if err != nil || bday.Before(time.Now()) {
		log.Fatal("Missing or wrong input dates")
	}
	return bday
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	hours := time.Until(target).Hours()
	days := hours/24 + 1
	return days
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
