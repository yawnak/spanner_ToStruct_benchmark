package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/spanner"
)

func main() {
	row, err := spanner.NewRow(rowColumnNames, rowColumnValues)
	if err != nil {
		log.Fatalln("error creating row:", err)
	}

	fmt.Println("Try convert a row first time...")

	var p1 Person
	err = row.ToStruct(&p1)
	if err != nil {
		log.Fatalln("error converting row to struct:", err)
	}
	fmt.Println("Person 1:", p1)

	fmt.Println("Try convert a row second time...")

	var p2 Person
	err = row.ToStruct(&p2)
	if err != nil {
		log.Fatalln("error converting row to struct second time:", err)
	}
	fmt.Println("Person 2:", p1)

}
