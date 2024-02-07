package checkout

import (
	"testing"

	"github.com/danieldavies99/checkout-kata/pricing"
)

// TestScan calls checkout.Scan with an item
// checking that the item is added to the
// scannedItems slice
func TestScan(t *testing.T) {
	// create TCheckout struct
	c := TCheckout{scannedItems: map[string]int{}}
	// create slice of test items
	items := []string{"A", "D", "F", "A"}

	// scan each item
	for _, item := range items {
		c.Scan(item)
	}

	// assert each item was counted correctly
	if c.scannedItems["A"] != 2 {
		t.Fatalf(`Scan() found %d A items, want %d, error`, c.scannedItems["A"], 2)
	}
	if c.scannedItems["D"] != 1 {
		t.Fatalf(`Scan() found %d D items, want %d, error`, c.scannedItems["C"], 1)
	}
	if c.scannedItems["F"] != 1 {
		t.Fatalf(`Scan() found %d F items, want %d, error`, c.scannedItems["F"], 2)
	}

	// probably unnecessary, but assert counts for items that are not scanned
	// = 0
	if c.scannedItems["B"] != 0 {
		t.Fatalf(`Scan() found %d B items, want %d, error`, c.scannedItems["B"], 0)
	}
}

// func TestGetTotalPriceNoMultiBuy asserts that the correct
// total checkout price is returned, only one of each item is scanned
// meaning that no multi-buy prices should be used
func TestGetTotalPriceNoMultiBuy(t *testing.T) {
	// create TPricing struct
	p := pricing.TPrices{Prices: map[string]pricing.TItemPrices{}}
	err := p.LoadFromJson("test_pricing.json")
	if err != nil {
		t.Fatalf(`LoadFromJson() return err:, %v`, err)
	}

	// create TCheckout struct
	c := TCheckout{
		scannedItems: map[string]int{},
		priceList:    p,
	}

	// create slice of test items
	items := []string{"A", "B", "C", "D"}

	// scan each item
	for _, item := range items {
		c.Scan(item)
	}

	totalPrice := c.GetTotalPrice()
	if totalPrice != 115 {
		t.Fatalf(`GetTotalPrice() got total: %v, want: %v`, totalPrice, 115)
	}
}

// func TestGetTotalPriceWithPerfectMultiBuy asserts that the correct
// total checkout price is returned, including multibuy deals where
// a perfect multiple of the multibuy cound is used.
//
// e.g. in our test data, SKU A has a multibuy deal of 3 for 130
// and SKU B has a multibuy deal of 2 for 45
// this test scans 3x A and 4x B
func TestGetTotalPriceWithPerfectMultiBuy(t *testing.T) {
	// create TPricing struct
	p := pricing.TPrices{Prices: map[string]pricing.TItemPrices{}}
	err := p.LoadFromJson("test_pricing.json")
	if err != nil {
		t.Fatalf(`LoadFromJson() return err:, %v`, err)
	}

	// create TCheckout struct
	c := TCheckout{
		scannedItems: map[string]int{},
		priceList:    p,
	}

	// create slice of test items
	items := []string{"A", "A", "A", "B", "B", "B", "B", "C", "D"}

	// scan each item
	for _, item := range items {
		c.Scan(item)
	}

	totalPrice := c.GetTotalPrice()
	if totalPrice != 255 {
		t.Fatalf(`GetTotalPrice() got total: %v, want: %v`, totalPrice, 255)
	}
}

// func TestGetTotalPriceWithImperfectMultiBuy asserts that the correct
// total checkout price is returned, including multibuy deals where
// a non-multiple of the multibuy cound is used.
//
// e.g. in our test data, SKU A has a multibuy deal of 3 for 130
// and SKU B has a multibuy deal of 2 for 45
// this test scans 4x A and 5x B
func TestGetTotalPriceWithImperfectMultiBuy(t *testing.T) {
	// create TPricing struct
	p := pricing.TPrices{Prices: map[string]pricing.TItemPrices{}}
	err := p.LoadFromJson("test_pricing.json")
	if err != nil {
		t.Fatalf(`LoadFromJson() return err:, %v`, err)
	}

	// create TCheckout struct
	c := TCheckout{
		scannedItems: map[string]int{},
		priceList:    p,
	}

	// create slice of test items
	items := []string{"A", "A", "A", "A", "B", "B", "B", "B", "B", "C", "D"}

	// scan each item
	for _, item := range items {
		c.Scan(item)
	}

	totalPrice := c.GetTotalPrice()
	if totalPrice != 335 {
		t.Fatalf(`GetTotalPrice() got total: %v, want: %v`, totalPrice, 335)
	}
}
