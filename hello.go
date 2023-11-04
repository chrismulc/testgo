package main

import (
	"encoding/json"
	"fmt"
	"os"
	"src/mypackage"
)

type Bill struct {
	Name   string
	Amount int
}

// test

func (b *Bill) setName(the_name string) {
	b.Name = the_name
}

func (b *Bill) setAmount(the_amount int) {
	b.Amount = the_amount
}

func (b Bill) getAmount() int {
	return b.Amount
}

func (b Bill) getName() string {
	return b.Name
}
func setNameOutside(b *Bill, the_name string) {
	b.Name = the_name
}

func saveBillToJSONFile(bill *Bill, filename string) error {
	// Marshal the Bill struct to JSON
	jsonData, err := json.Marshal(bill)
	if err != nil {
		return err
	}

	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	fmt.Printf("Bill data saved to %s\n", filename)
	return nil
}

func readJsonFile(filename string) (*Bill, error) {
	// Open the JSON file for reading
	the_file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer the_file.Close()

	// Create a Bill struct to store the JSON data
	var internalBill Bill

	// Create a JSON decoder
	decoder := json.NewDecoder(the_file)

	// Decode the JSON data into the Bill struct
	if err := decoder.Decode(&internalBill); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	return &internalBill, nil
}

func my_func() {
	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')

	// bill1.setName(text)

	// setNameOutside(&bill1, "shit again")

	// fmt.Println(bill1.getName())
	// bill1.setAmount(77)

	// bills = append(bills, bill1)
	// bills = append(bills, bill2)

	// test := bills[1]
	// shit2 := test.getAmount()
	// fmt.Println(shit2)
	// fmt.Println("amount2: ", test.getName())
	// fmt.Println("text: ", text)

	// billtest, _ := readJsonFile("bill1.json")
	// fmt.Printf("Bill: %+v\n", billtest)
}

func main() {

	//var bills []Bill

	//bill1 := Bill{
	//	Name:   "chris",
	//	Amount: 999,
	//}

	//bill2 := Bill{
	//	Name:   "danni",
	//	Amount: 100,
	//}

	/*
		err := saveBillToJSONFile(&bill1, "bill.json")

		if err != nil {
			fmt.Println("Error:", err)
		}
	*/

	var my_test mypackage.Faker
	my_test.Name = "Faker"
	

	billtest, _ := readJsonFile("bill.json")
	fmt.Printf("Bill : %+v\n", billtest)

}
