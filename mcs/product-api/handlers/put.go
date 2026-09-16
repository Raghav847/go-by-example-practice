package handlers

import (
	"errors"
	"net/http"

	"github.com/Raghav847/go-by-example-practice/mcs/product-api/data"
)

func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("[DEBUG] Updating product: %+v\n", prod)

	err := data.UpdateProduct(prod)
	if errors.Is(err, data.ErrProductNotFound) {
		p.l.Println("[ERROR] product not found", err)
		http.Error(w, "product not found", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
