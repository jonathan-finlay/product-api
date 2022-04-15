package data

import (
	"encoding/json"
	"io"
	"time"
	"fmt"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lastProduct := productList[len(productList) -1]
	return lastProduct.ID + 1
}

func UpdatedProduct(id int, p *Product) error {
	_, position, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[position] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "asdf123",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Small and strong coffee w/o milk",
		Price:       1.99,
		SKU:         "qwe098",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
