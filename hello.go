package main

import "fmt"

type bill struct {
	name   string
	amount int
}

func main() {

	bill1 := bill{
		name:   "chris",
		amount: 99,
	}
	fmt.Println(bill1.name)

}
