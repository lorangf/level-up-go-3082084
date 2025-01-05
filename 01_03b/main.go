package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("Could not read the file %v", path)
	}
	var entries []raffleEntry
	if err := json.Unmarshal(contentBytes, &entries); err != nil {
		log.Panicf("Could not unmarshall the data, %v", err)
	}
	return entries
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	generator := rand.New(rand.NewSource(time.Now().Unix()))
	wi := generator.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
