package main

type Person struct {
	ID        int64   `spanner:"id"`
	Name      string  `spanner:"name"`
	Age       int64   `spanner:"age"`
	Weight    float64 `spanner:"weight"`
	IsMarried bool    `spanner:"is_married"`
}

func (p *Person) Columns() []interface{} {
	return []interface{}{
		&p.ID,
		&p.Name,
		&p.Age,
		&p.Weight,
		&p.IsMarried,
	}
}

var (
	rowColumnNames  = []string{"id", "name", "age", "weight", "is_married"}
	rowColumnValues = []interface{}{1, "John", 20, 50.5, true}
)
