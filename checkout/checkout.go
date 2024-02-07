package checkout

type ICheckout interface {
	Scan(string)
	GetTotalPrice() int
}

type TCheckout struct {
	scannedItems []string
}

// Scan inserts a given item into the scannedItems
// property, later we will use the scannedItems
// list to work out the final price
func (c *TCheckout) Scan(item string) {

	// should I check if scanned item exists?
	c.scannedItems = append(c.scannedItems, item)
}
