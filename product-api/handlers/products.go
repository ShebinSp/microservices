package handlers

import (
	//"encoding/json"
	"log"
	"microServices/product-api/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// Catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)	
}
func (p * Products) getProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle Get Products")
	// below the json.Marshal() is used to send small amount of values(small resposes)
	//d, err := json.Marshal(lp)
	// if err != nil {
	// 	http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	// }
	//rw.Write(d)

	// Here we are using function to encode JSON.
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	} 
}

func (p *Products) addProduct (rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle Post Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Product : %#v -added",prod)
	data.AddProducts(prod)
}