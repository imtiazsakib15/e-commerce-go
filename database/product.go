package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var products []Product

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

	products = append(
		products, product1, product2)
}

func Store(product Product) Product {
	product.ID = len(products) + 1
	products = append(products, product)
	return product
}

func List() []Product {
	return products
}

func Get(productId int) *Product {
	for _, product := range products {
		if product.ID == productId {
			return &product
		}
	}
	return nil
}

func Update(productID int, product Product) *Product {
	for idx, pdt := range products {
		if pdt.ID == productID {
			product.ID = productID
			products[idx] = product
			return &product
		}
	}
	return nil
}

func Delete(productID int) *Product {
	for idx, pdt := range products {
		if pdt.ID == productID {
			products[idx] = Product{}
			return &pdt
		}
	}
	return nil
}