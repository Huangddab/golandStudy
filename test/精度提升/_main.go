package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func init() {
	// 引入第三方包解决精度提升问题
	var num1 float64 = 3.1
	var num2 float64 = 4.3
	d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	fmt.Println(d1)
}
