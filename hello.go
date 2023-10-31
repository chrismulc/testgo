package main

import (
	"bufio"
	"fmt"
	"os"
)

type bill struct {
	name   string
	amount int
}

// test

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

	var bills []bill

	i := 99

	bill1 := bill{
		name:   "chris",
		amount: 999,
	}

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	bill1.setName(text)

	// bill1.setName("test")
	setNameOutside(&bill1, "shit again")

	fmt.Println(bill1.getName())
	bill1.setAmount(i)
	bill1.setAmount(77)

	bills = append(bills, bill1)
	test := bills[0]
	fmt.Println("amount: ", test.getAmount())
}
