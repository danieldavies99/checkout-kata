package checkout

import (
	"testing"
)

// TestScan calls checkout.Scan with an item
// checking that the item is added to the
// scannedItems slice
func TestScan(t *testing.T) {
	// create TCheckout struct
	c := TCheckout{scannedItems: []string{}}
	// create slice of test items
	items := []string{"A", "D", "F", "A"}

	// scan each item
	for _, item := range items {
		c.Scan(item)
	}

	// assert each item was scanned
	for i, scannedItem := range c.scannedItems {

		if scannedItem != items[i] {
			t.Fatalf(`Scan() found item, %s, want %s, error`, scannedItem, items[i])
		}
	}
}
