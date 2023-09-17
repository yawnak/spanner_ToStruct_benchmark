package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/spanner"
)

func main() {
	// Create a row from column names and values.
	row, err := spanner.NewRow(rowColumnNames, rowColumnValues)
	if err != nil {
		log.Fatalln("error creating row:", err)
	}

	fmt.Println("Try convert a row first time...")
	// Convert the row to a struct.
	var p1 Person
	err = row.ToStruct(&p1)
	if err != nil {
		log.Fatalln("error converting row to struct:", err)
	}
	fmt.Println("Person 1:", p1)

	fmt.Println("Try convert a row second time...")
	// Convert the row to a struct.
	var p2 Person
	err = row.ToStruct(&p2)
	if err != nil {
		log.Fatalln("error converting row to struct second time:", err)
	}
	fmt.Println("Person 2:", p1)
}
