package main

import (
	"testing"

	"cloud.google.com/go/spanner"
)

func BenchmarkRowToStruct(b *testing.B) {
	row, err := spanner.NewRow(rowColumnNames, rowColumnValues)
	if err != nil {
		b.Fatal("error creating row:", err)
	}
	var p Person

	// Reset the timer to ignore initialization overhead.
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		row.ToStruct(&p)
	}
}

func BenchmarkRowColumns(b *testing.B) {
	row, err := spanner.NewRow(rowColumnNames, rowColumnValues)
	if err != nil {
		b.Fatal("error creating row:", err)
	}
	var p Person

	// Reset the timer to ignore initialization overhead.
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		row.Columns(
			&p.ID,
			&p.Name,
			&p.Age,
			&p.Weight,
			&p.IsMarried,
		)
	}
}
