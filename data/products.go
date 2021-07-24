package data

import (
	"encoding/json"
	"io"
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

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Coffee",
		Price:       10.0,
		SKU:         "Something",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          2,
		Name:        "Mocca",
		Description: "Coffee",
		Price:       12.0,
		SKU:         "Something",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
}
