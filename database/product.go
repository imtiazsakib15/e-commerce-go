package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var Products []Product

func init() {
	product1 :=
		Product{
			ID:          1,
			Title:       "Mango",
			Description: "1st Fruit",
			Price:       22,
		}
	product2 :=
		Product{
			ID:          2,
			Title:       "Banana",
			Description: "2nd Fruit",
			Price:       12,
		}

	Products = append(
		Products, product1, product2)
}