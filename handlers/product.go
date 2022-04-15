package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"../data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Test 0")
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`products/(0-9}+)`)
		captureGroup := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(captureGroup) != 1 {
			p.l.Println("Test 1")
			http.Error(w, "Invalid URI, more than one id", http.StatusBadRequest)
			return
		}

		if len(captureGroup[0]) != 2 {
			p.l.Println("Test 2")
			http.Error(w, "Invalid URI, more than one capture group", http.StatusBadRequest)
			return
		}

		idString := captureGroup[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Test 3")
			http.Error(w, "Invalid URI, unable to convert number", http.StatusBadRequest)
		}

		p.l.Println("Test 4")
		p.updateProducts(id, w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	newProduct := &data.Product{}
	err := newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadGateway)
	}

	p.l.Printf("Product: %#v", newProduct)
	data.AddProduct(newProduct)
}

func (p *Products) updateProducts(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	updatedProduct := &data.Product{}
	err := updatedProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadGateway)
	}

	err = data.UpdatedProduct(id, updatedProduct)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
}
