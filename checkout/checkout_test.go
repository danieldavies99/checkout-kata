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

func TestGetTotalPriceNoMultiBuy(t *testing.T) {
	// create TPricing struct
	p := pricing.TPrices{}
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

func TestGetTotalPriceWithPerfectMultiBuy(t *testing.T) {
	// create TPricing struct
	p := pricing.TPrices{}
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

func TestGetTotalPriceWithImperfectMultiBuy(t *testing.T) {
	// create TPricing struct
	p := pricing.TPrices{}
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
