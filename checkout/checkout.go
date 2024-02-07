package checkout

import (
	"math"

	"github.com/danieldavies99/checkout-kata/pricing"
)

type ICheckout interface {
	Scan(string)
	GetTotalPrice() int
}

type TCheckout struct {
	scannedItems map[string]int // map[itemSku]quantityScanned
	priceList    pricing.TPrices
}

// Scan inserts a given item into the scannedItems
// map
func (c *TCheckout) Scan(item string) {
	// should I check if scanned item exists?
	c.scannedItems[item] += 1
}

// GetTotalPrice returns the total price of scanned items
// based on the provided priceList
func (c TCheckout) GetTotalPrice() int {
	// iterate over every item in pricelist and see if we
	// have scanned any items of that type, wouldn't be
	// very efficient if we had a large pricelist
	total := 0
	for _, itemPrices := range c.priceList.Prices {
		// we didn't scan any of these items
		// so skip
		numScanned := c.scannedItems[itemPrices.Sku]
		if numScanned == 0 {
			continue
		}

		// there is a multibuy discount
		if itemPrices.MultiBuyCount != nil && itemPrices.MultiBuyPrice != nil {
			// get count of multibuy total
			// e.g. if discount is 3 for 10
			// and we've scanned 10 items
			// multibuy total = 3
			// idividualPriceTotal = 1
			multiBuyTotal := math.Floor(float64(numScanned) / float64(*itemPrices.MultiBuyCount))
			individualPriceTotal := numScanned % *itemPrices.MultiBuyCount
			total += int(multiBuyTotal) * *itemPrices.MultiBuyPrice
			total += individualPriceTotal * *&itemPrices.UnitPrice
			continue
		}

		// there is no multibuy discount
		total += numScanned * itemPrices.UnitPrice
	}

	return total
}
