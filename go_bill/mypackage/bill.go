package mypackage

import (
	"encoding/json"
	"fmt"
	"os"
)




type Faker struct {
	Name string
}

func (f Faker) getName() string {
	return f.Name
}

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

func ReadJsonFile(filename string) (*Bill, error) {
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

