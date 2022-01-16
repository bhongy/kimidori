package data

import (
	"errors"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func GetProducts() []*Product {
	return productLits
}

func AddProduct(p *Product) {
	p.ID = nextID()
	t := time.Now().UTC().String()
	p.CreatedOn = t
	p.UpdatedOn = t
	productLits = append(productLits, p)
}

func nextID() int {
	return len(productLits) + 1
}

var ErrProductNotFound = errors.New("Product not found.")

func UpdateProduct(p *Product) error {
	_, i, err := findProductByID(p.ID)
	if err != nil {
		return err
	}
	productLits[i] = p
	return nil
}

func findProductByID(id int) (*Product, int, error) {
	for i, p := range productLits {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var productLits = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
