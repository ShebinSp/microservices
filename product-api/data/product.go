package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreatedOn string `json:"-"` // the field is ommitted in json output
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Encoding go value to JSON format. It is helpful when large amount of data is need to transfered or when have to 
// store the data(when using microservices)
type Products []*Product
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func GetProducts()  Products {
	return productList
}

func AddProducts(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)

}

func getNextID() int {
	lastProd := productList[len(productList)-1]
	return lastProd.ID + 1
}

var productList = []*Product {
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffe",
		Price: 150,
		SKU: "lat123",
		CreatedOn: time.Now().String(),
		UpdatedOn: time.Now().String(),
		
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffer without milk",
		Price: 100,
		SKU: "esp123",
		CreatedOn: time.Now().String(),
		UpdatedOn: time.Now().String(),
	},
}