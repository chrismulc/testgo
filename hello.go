package main

import "fmt"

type bill struct {
	name   string
	amount int
}

func (b *bill) setName(the_name string) {
	b.name = the_name
}

func (b *bill) setAmount(the_amount int) {
	b.amount = the_amount
}

func (b bill) getAmount() int {
	return b.amount
}

func (b bill) getName() string {
	return b.name
}
func setNameOutside(b *bill, the_name string) {
	b.name = the_name
}

func main() {

	i := 99

	bill1 := bill{
		name:   "chris",
		amount: 99,
	}

	bill1.setName("test")
	setNameOutside(&bill1, "shit again")
	fmt.Println(bill1.getName())
	bill1.setAmount(i)
	bill1.setAmount(77)
	fmt.Println("amount: ", bill1.getAmount())
}
