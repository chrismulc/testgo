package main

import "fmt"

type bill struct {
	name   string
	amount int
}

func (b *bill) set_name(the_name string) {
	b.name = the_name
}

func set_name_outside(b *bill, the_name string) {
	b.name = the_name
}

func main() {

	bill1 := bill{
		name:   "chris",
		amount: 99,
	}

	bill1.set_name("test")
	set_name_outside(&bill1, "shit")
	fmt.Println(bill1.name)

}
