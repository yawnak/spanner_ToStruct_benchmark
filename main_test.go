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

func BenchmarkRowToStructBig(b *testing.B) {
	row, err := spanner.NewRow(bigRowColumnNames, bigRowColumnValues)
	if err != nil {
		b.Fatal("error creating row:", err)
	}
	var p BigPerson

	// Reset the timer to ignore initialization overhead.
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		row.ToStruct(&p)
	}
}

func BenchmarkRowColumnsBig(b *testing.B) {
	row, err := spanner.NewRow(bigRowColumnNames, bigRowColumnValues)
	if err != nil {
		b.Fatal("error creating row:", err)
	}
	var p BigPerson

	// Reset the timer to ignore initialization overhead.
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		row.Columns(
			&p.ID,
			&p.Name,
			&p.Age,
			&p.Weight,
			&p.IsMarried,
			&p.Extra,
			&p.Driver,
			&p.Height,
			&p.Width,
			&p.FavoriteFood,
			&p.FavoriteColor,
			&p.FavoriteNumber,
			&p.FavoriteAnimal,
			&p.FavoriteMovie,
			&p.FavoriteBook,
			&p.FavoriteSong,
			&p.FavoriteGame,
			&p.FavoriteSport,
			&p.FavoriteTeam,
			&p.FavoriteBand,
			&p.FavoriteDrink,
			&p.FavoritePlace,
			&p.FavoriteCar,
			&p.FavoriteShoe,
			&p.FavoriteCandy,
			&p.FavoriteIceCream,
			&p.FavoriteFruit,
			&p.FavoriteVegetable,
			&p.FavoriteFlower,
			&p.FavoriteTree,
			&p.FavoritePlanet,
			&p.FavoriteElement,
		)
	}
}
