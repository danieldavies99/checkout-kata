package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/danieldavies99/checkout-kata/checkout"
	"github.com/danieldavies99/checkout-kata/pricing"
)

func main() {
	// create TPrices struct
	p := pricing.TPrices{Prices: map[string]pricing.TItemPrices{}}
	err := p.LoadFromJson("pricing.json")
	if err != nil {
		log.Fatalf(`LoadFromJson() return err:, %v`, err)
	}

	// create TCheckout struct
	c := checkout.TCheckout{
		ScannedItems: map[string]int{},
		PriceList:    p,
	}
	// create slice of test items
	items := []string{}
	// add a random number (0-20) of each item A, B, C, D
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 20
	for i := 1; i <= rand.Intn(max-min)+min; i++ {
		items = append(items, "A")
	}
	for i := 1; i <= rand.Intn(max-min)+min; i++ {
		items = append(items, "B")
	}
	for i := 1; i <= rand.Intn(max-min)+min; i++ {
		items = append(items, "C")
	}
	for i := 1; i <= rand.Intn(max-min)+min; i++ {
		items = append(items, "D")
	}

	// scan each item
	for _, item := range items {
		err = c.Scan(item)
		if err != nil {
			log.Fatalf(`Scan() returned error: %v`, err)
		}
	}

	total := c.GetTotalPrice()

	log.Printf(
	`The total price for
A:%v
B:%v
C:%v
D:%v
Is %v`,
		c.ScannedItems["A"],
		c.ScannedItems["B"],
		c.ScannedItems["C"],
		c.ScannedItems["D"],
		total,
	)
	os.Exit(0)
}
