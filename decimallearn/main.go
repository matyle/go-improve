package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Dig struct {
	num decimal.Decimal
}

func main() {
	d := new(Dig)
	d.num, _ = decimal.NewFromString("0")
	fmt.Println(d.num)
	fmt.Println(d.num.Equal(decimal.Zero))
}
