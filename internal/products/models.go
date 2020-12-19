package products

import (
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID               string
	Name             string
	Price            decimal.Decimal
	PriceUpdateCount uint32
	LastModified     time.Time
}

type Paging struct {
	Offset uint32
	Limit  uint32
}

type Sorting struct {
	Ascending bool
	SortBy    string
}

var fieldsToSortBy = []string{
	"name",
	"price",
	"priceUpdateCount",
	"lastModified",
}

func (s Sorting) Validate() error {
	for _, field := range fieldsToSortBy {
		if strings.EqualFold(s.SortBy, field) {
			return nil
		}
	}
	return fmt.Errorf("Validate: can not sort products by field: %s", s.SortBy)
}
