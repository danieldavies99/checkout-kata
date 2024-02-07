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
	Prices []TItemPrices `json:"items"`
}

// LoadFromJson takes a path to a json file
// containing pricing information and parses
// it into my structs as defined above
func (c *TPrices) LoadFromJson(path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Failed to open json file, %v", err)
	}

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("Failed to read bytes json file, %v", err)
	}

	json.Unmarshal(bytes, &c)

	defer jsonFile.Close()
	return nil
}
