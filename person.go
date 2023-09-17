package main

type Person struct {
	ID        int64   `spanner:"id"`
	Name      string  `spanner:"name"`
	Age       int64   `spanner:"age"`
	Weight    float64 `spanner:"weight"`
	IsMarried bool    `spanner:"is_married"`
}
