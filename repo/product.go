package repo

import "errors"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductRepo interface {
	Create(product Product) (*Product, error)
	List() (*[]Product, error)
	GetById(productId int) (*Product, error)
	Update(productId int, product Product) (*Product, error)
	Delete(productId int) (*Product, error)
}

type productRepo struct {
	products []Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(product Product) (*Product, error) {
	product.ID = len(r.products) + 1
	r.products = append(r.products, product)
	return &r.products[len(r.products)-1], nil
}
func (r *productRepo) List() (*[]Product, error) {
	return &r.products, nil
}
func (r *productRepo) GetById(productId int) (*Product, error) {
	for idx, product := range r.products {
		if product.ID == productId {
			return &r.products[idx], nil
		}
	}
	return nil, errors.New("Product not found")
}
func (r *productRepo) Update(productId int, product Product) (*Product, error) {
	for idx, pdt := range r.products {
		if pdt.ID == productId {
			product.ID = productId
			r.products[idx] = product
			return &r.products[idx], nil
		}
	}
	return nil, errors.New("Product not found")
}
func (r *productRepo) Delete(productId int) (*Product, error) {
	for idx, pdt := range r.products {
		if pdt.ID == productId {
			deleted := r.products[idx]
			r.products = append(r.products[:idx], r.products[idx+1:]...)
			return &deleted, nil
		}
	}
	return nil, errors.New("Product not found")
}

func generateInitialProducts(r *productRepo) {
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

	r.products = append(
		r.products, product1, product2)
}
