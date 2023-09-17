package main

type BigPerson struct {
	ID                int64   `spanner:"id"`
	Name              string  `spanner:"name"`
	Age               int64   `spanner:"age"`
	Weight            float64 `spanner:"weight"`
	IsMarried         bool    `spanner:"is_married"`
	Extra             string  `spanner:"extra"`
	Driver            string  `spanner:"driver"`
	Height            int64   `spanner:"height"`
	Width             int64   `spanner:"width"`
	FavoriteFood      string  `spanner:"favorite_food"`
	FavoriteColor     string  `spanner:"favorite_color"`
	FavoriteNumber    int64   `spanner:"favorite_number"`
	FavoriteAnimal    string  `spanner:"favorite_animal"`
	FavoriteMovie     string  `spanner:"favorite_movie"`
	FavoriteBook      string  `spanner:"favorite_book"`
	FavoriteSong      string  `spanner:"favorite_song"`
	FavoriteGame      string  `spanner:"favorite_game"`
	FavoriteSport     string  `spanner:"favorite_sport"`
	FavoriteTeam      string  `spanner:"favorite_team"`
	FavoriteBand      string  `spanner:"favorite_band"`
	FavoriteDrink     string  `spanner:"favorite_drink"`
	FavoritePlace     string  `spanner:"favorite_place"`
	FavoriteCar       string  `spanner:"favorite_car"`
	FavoriteShoe      string  `spanner:"favorite_shoe"`
	FavoriteCandy     string  `spanner:"favorite_candy"`
	FavoriteIceCream  string  `spanner:"favorite_ice_cream"`
	FavoriteFruit     string  `spanner:"favorite_fruit"`
	FavoriteVegetable string  `spanner:"favorite_vegetable"`
	FavoriteFlower    string  `spanner:"favorite_flower"`
	FavoriteTree      string  `spanner:"favorite_tree"`
	FavoritePlanet    string  `spanner:"favorite_planet"`
	FavoriteElement   string  `spanner:"favorite_element"`
}

var (
	//column names for row bigperson
	bigRowColumnNames = []string{
		"id", "name", "age", "weight", "is_married",
		"extra", "driver", "height", "width", "favorite_food",
		"favorite_color", "favorite_number", "favorite_animal",
		"favorite_movie", "favorite_book", "favorite_song", "favorite_game",
		"favorite_sport", "favorite_team", "favorite_band", "favorite_drink",
		"favorite_place", "favorite_car", "favorite_shoe", "favorite_candy",
		"favorite_ice_cream", "favorite_fruit", "favorite_vegetable",
		"favorite_flower", "favorite_tree", "favorite_planet",
		"favorite_element",
	}
	bigRowColumnValues = []interface{}{
		1, "John", 20, 50.5, true,
		"extra", "driver", 1, 1, "favorite_food",
		"favorite_color", 1, "favorite_animal",
		"favorite_movie", "favorite_book", "favorite_song", "favorite_game",
		"favorite_sport", "favorite_team", "favorite_band", "favorite_drink",
		"favorite_place", "favorite_car", "favorite_shoe", "favorite_candy",
		"favorite_ice_cream", "favorite_fruit", "favorite_vegetable",
		"favorite_flower", "favorite_tree", "favorite_planet",
		"favorite_element",
	}
)
