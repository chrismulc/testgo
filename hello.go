package main

import "fmt"

type bill struct {
	name   string
	amount int
}

func (b *bill) setName(the_name string) {
	b.name = the_name
}

func (b *bill) getName() string {
	return b.name
}
func setNameOutside(b *bill, the_name string) {
	b.name = the_name
}

func main() {

	bill1 := bill{
		name:   "chris",
		amount: 99,
	}

	bill1.setName("test")
	setNameOutside(&bill1, "shit again")
	fmt.Println(bill1.getName())

}
