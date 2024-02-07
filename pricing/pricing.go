package pricing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type IPriceMap interface {
	LoadFromJson(string) error
	GetPricingForItem(string) TItemPrices
}

type TItemPrices struct {
	Sku           string `json:"sku"`
	UnitPrice     int    `json:"unit_price"`     // price if bought individually
	MultiBuyPrice *int   `json:"multibuy_price"` // price if {multiBuyCount} bought
	MultiBuyCount *int   `json:"multibuy_quantity"`
}

type TPrices struct {
	Prices map[string]TItemPrices
}

// LoadFromJson takes a path to a json file
// containing pricing information and parses
// it into the structs as defined above
// I refactored this to output a map instead of
// my original slice, this is because I can address
// the pricing directly via Prices[item] rather than
// having to do a search
func (c *TPrices) LoadFromJson(path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Failed to open json file, %v", err)
	}

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("Failed to read bytes json file, %v", err)
	}

	var t struct {
		Prices []TItemPrices `json:"items"`
	}

	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal json, %v", err)
	}

	for _, v := range t.Prices {
		c.Prices[v.Sku] = v
	}

	defer jsonFile.Close()
	return nil
}
