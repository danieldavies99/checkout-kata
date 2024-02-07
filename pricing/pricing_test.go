package pricing

import (
	"testing"
)

// TestLoadFromJson calls pricing.LoadFromJson
// checking that the expected test pricing
// data is loaded
func TestLoadFromJson(t *testing.T) {
	p := TPrices{map[string]TItemPrices{}}
	err := p.LoadFromJson("test_pricing.json")
	if err != nil {
		t.Fatalf(`LoadFromJson() return err:, %v`, err)
	}

	// Check first item
	// I used 4 ifs rather than comparing two complete structs because
	// I am using pointers for MultiBuyPrice and MultiBuyCount (because they might be nil) which
	// would break the comparision as Go would try to compare the memory addresses
	// rather than the values pointed to
	if p.Prices["A"].Sku != "A" {
		t.Fatalf(`LoadFromJson() found Sku: %v, want: %v`, p.Prices["A"].Sku, "A")
	}
	if p.Prices["A"].UnitPrice != 50 {
		t.Fatalf(`LoadFromJson() found UnitPrice: %v, want: %v`, p.Prices["A"].UnitPrice, 50)
	}
	if *p.Prices["A"].MultiBuyPrice != 130 {
		t.Fatalf(`LoadFromJson() found MultiBuyPrice: %v, want: %v`, *p.Prices["A"].MultiBuyPrice, 130)
	}
	if *p.Prices["A"].MultiBuyCount != 3 {
		t.Fatalf(`LoadFromJson() found MultiBuyPrice: %v, want: %v`, *p.Prices["A"].MultiBuyCount, 3)
	}

	// Check last item:
	if p.Prices["D"].Sku != "D" {
		t.Fatalf(`LoadFromJson() found Sku: %v, want: %v`, p.Prices["D"].Sku, "D")
	}
	if p.Prices["D"].UnitPrice != 15 {
		t.Fatalf(`LoadFromJson() found UnitPrice: %v, want: %v`, p.Prices["D"].UnitPrice, 15)
	}
	if p.Prices["D"].MultiBuyPrice != nil {
		t.Fatalf(`LoadFromJson() found MultiBuyPrice: %v, want: %v`, p.Prices["D"].MultiBuyPrice, "Nil pointer")
	}
	if p.Prices["D"].MultiBuyCount != nil {
		t.Fatalf(`LoadFromJson() found MultiBuyPrice: %v, want: %v`, p.Prices["D"].MultiBuyCount, "Nil pointer")
	}
}
