package main

import (
	"fmt"
	"go_bill/mypackage"
)

func main() {

	// this is in test
	var my_test mypackage.Faker
	my_test.Name = "Chris"
	var crap *mypackage.Bill
	crap, _ = mypackage.ReadJsonFile("bill.json")
	// only in test
	// also only in test

	fmt.Printf("%s", my_test.Name)
	billtest, _ := mypackage.ReadJsonFile("bill.json")
	fmt.Printf("Bill : %+v\n", billtest)
	fmt.Printf("Crap : %+v\n", crap)
	// vscode
}
