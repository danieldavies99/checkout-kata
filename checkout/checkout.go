package checkout

import (
	"fmt"
	"math"

	"github.com/danieldavies99/checkout-kata/pricing"
)

type ICheckout interface {
	Scan(string) error
	GetTotalPrice() int
}

type TCheckout struct {
	ScannedItems map[string]int // map[itemSku]quantityScanned
	PriceList    pricing.TPrices
}

// Scan inserts a given item into the scannedItems
// map
func (c *TCheckout) Scan(item string) error {
	// check if price exists
	if _, ok := c.PriceList.Prices[item]; !ok {
		return fmt.Errorf("Attempted to scan item that doesn't exist in pricelist")
	}
	c.ScannedItems[item] += 1
	return nil
}

// GetTotalPrice returns the total price of scanned items
// based on the provided priceList
func (c TCheckout) GetTotalPrice() int {
	// iterate over each value in the scanned items map
	// this is better than my original approach of iterating through
	// all items in the pricelist because this will scale with large
	// pricelists
	total := 0
	for item, scanCount := range c.ScannedItems {
		itemPricing := c.PriceList.Prices[item]
		// there is a multibuy discount
		if itemPricing.MultiBuyCount != nil && itemPricing.MultiBuyPrice != nil {
			// get count of multibuy total
			// e.g. if discount is 3 for 10
			// and we've scanned 10 items
			// multiBuyTotal = 3
			// idividualPriceTotal = 1
			multiBuyTotal := math.Floor(float64(scanCount) / float64(*itemPricing.MultiBuyCount))
			individualPriceTotal := scanCount % *itemPricing.MultiBuyCount
			total += int(multiBuyTotal) * *itemPricing.MultiBuyPrice
			total += individualPriceTotal * itemPricing.UnitPrice
			continue
		}

		// there is no multibuy discount
		total += scanCount * itemPricing.UnitPrice
	}

	return total
}
