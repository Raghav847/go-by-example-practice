package handlers

import (
	"net/http"

	"github.com/Raghav847/go-by-example-practice/mcs/product-api/data"
)

func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("[DEBUG] Inserting product: %+v\n", prod)
	data.AddProduct(prod)
}
