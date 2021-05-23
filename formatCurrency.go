// Created by: Maryam Nona
// Created on: May 2021
//
// This program formats currency

package main

import (
	"fmt"
	"github.com/leekchan/accounting"
)

func main() {
	// this function displays currency
	accountingFormater := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(accountingFormater.FormatMoney(1058.2374))
}