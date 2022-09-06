package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Dig struct {
	num decimal.Decimal
}

func main() {
	decimal.DivisionPrecision = 28
	nums := "0.0036529728291999"
	d, err := decimal.NewFromString(nums)
	if err != nil {
		fmt.Println(err)
	}
	d1 := d
	fmt.Println(d1)
}
