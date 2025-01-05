package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
)

const path = "items.json"

// SaleItem represents the item part of the big sale.
type SaleItem struct {
	Name           string  `json:"name"`
	OriginalPrice  float64 `json:"originalPrice"`
	ReducedPrice   float64 `json:"reducedPrice"`
	SalePercentage float64
}

// type BySalePercentage []SaleItem

// func (a BySalePercentage) Len() int           { return len(a) }
// func (a BySalePercentage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a BySalePercentage) Less(i, j int) bool { return a[i].SalePercentage < a[j].SalePercentage }

// matchSales adds the sales procentage of the item
// and sorts the array accordingly.
func matchSales(budget float64, items []SaleItem) []SaleItem {
	itemsInBudget := make([]SaleItem, 0)
	for i := 0; i < len(items); i++ {
		if budget >= items[i].ReducedPrice {
			items[i].SalePercentage = (items[i].OriginalPrice - items[i].ReducedPrice) / items[i].OriginalPrice * 100.0
			itemsInBudget = append(itemsInBudget, items[i])
		}
	}

	sort.Slice(itemsInBudget, func(i, j int) bool {
		return itemsInBudget[i].SalePercentage > itemsInBudget[j].SalePercentage
	})

	return itemsInBudget
}

func main() {
	budget := flag.Float64("budget", 0.0,
		"The max budget you want to shop with.")
	flag.Parse()
	items := importData()
	matchedItems := matchSales(*budget, items)
	printItems(matchedItems)
}

// printItems prints the items and their sales.
func printItems(items []SaleItem) {
	log.Println("The BIG sale has started with our amazing offers!")
	if len(items) == 0 {
		log.Println("No items found.:( Try increasing your budget.")
	}
	for i, r := range items {
		log.Printf("[%d]:%s is %.2f OFF! Get it now for JUST %.2f!\n",
			i, r.Name, r.SalePercentage, r.ReducedPrice)
	}
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []SaleItem {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []SaleItem
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
