package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	words := strings.Split(msg, " ")
	for _, w := range words {
		var text []byte
		for i := 0; i < len(w); i++ {
			for j := 0; j <= i; j++ {
				text = append(text, w[i])
			}
		}
		print(string(text))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
