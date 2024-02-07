package pricing

import (
	"testing"
)

// TestLoadFromJson calls pricing.LoadFromJson
// checking that the expected test pricing
// data is loaded
func TestLoadFromJson(t *testing.T) {
	p := TPrices{}
	err := p.LoadFromJson("test_pricing.json")
	if err != nil {
		t.Fatalf(`LoadFromJson() return err:, %v`, err)
	}

	// Check first item
	// I used 4 ifs rather than comparing two complete structs because
	// I am using pointers for MultiBuyPrice and MultiBuyCount which
	// would break the comparision as go would try to compare the memory addresses
	// rather than the values pointed to
	if p.Prices[0].Sku != "A" {
		t.Fatalf(`LoadFromJson() found Sku: %v, want: %v`, p.Prices[0].Sku, "A")
	}
	if p.Prices[0].UnitPrice != 50 {
		t.Fatalf(`LoadFromJson() found UnitPrice: %v, want: %v`, p.Prices[0].UnitPrice, 50)
	}
	if *p.Prices[0].MultiBuyPrice != 130 {
		t.Fatalf(`LoadFromJson() found MultiBuyPrice: %v, want: %v`, *p.Prices[0].MultiBuyPrice, 130)
	}
	if *p.Prices[0].MultiBuyCount != 3 {
		t.Fatalf(`LoadFromJson() found MultiBuyPrice: %v, want: %v`, *p.Prices[0].MultiBuyCount, 3)
	}

	// Check last item:
	lastItemIndex := len(p.Prices) - 1
	if p.Prices[lastItemIndex].Sku != "D" {
		t.Fatalf(`LoadFromJson() found Sku: %v, want: %v`, p.Prices[0].Sku, "D")
	}
	if p.Prices[lastItemIndex].UnitPrice != 15 {
		t.Fatalf(`LoadFromJson() found UnitPrice: %v, want: %v`, p.Prices[0].UnitPrice, 15)
	}
	if p.Prices[lastItemIndex].MultiBuyPrice != nil {
		t.Fatalf(`LoadFromJson() found MultiBuyPrice: %v, want: %v`, p.Prices[0].MultiBuyPrice, "Nil pointer")
	}
	if p.Prices[lastItemIndex].MultiBuyCount != nil {
		t.Fatalf(`LoadFromJson() found MultiBuyPrice: %v, want: %v`, p.Prices[0].MultiBuyCount, "Nil pointer")
	}
}
