package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Raghav847/go-by-example-practice/mcs/product-api/data"
)

type KeyProduct struct{}

type Products struct {
	l *log.Logger
	v *data.Validation
}

func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func getProductID(r *http.Request) int {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(err)
	}

	return id
}

//func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
//	p.l.Println("Handle POST Product")
//
//	prod, ok := r.Context().Value(KeyProduct{}).(*data.Product)
//	if !ok {
//		http.Error(w, "validated product missing", http.StatusInternalServerError)
//		return
//	}
//
//	data.AddProduct(prod)
//}
//
//func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
//	p.l.Println("Handle PUT Products")
//
//	id, err := strconv.Atoi(r.PathValue("id"))
//	if err != nil {
//		http.Error(w, "invalid id", http.StatusBadRequest)
//		return
//	}
//
//	prod, ok := r.Context().Value(KeyProduct{}).(*data.Product)
//	if !ok {
//		http.Error(w, "validated product missing", http.StatusInternalServerError)
//		return
//	}
//
//	err = data.UpdateProduct(id, prod)
//	if errors.Is(err, data.ErrProductNotFound) {
//		http.Error(w, "product not found", http.StatusNotFound)
//		return
//	}
//
//	if err != nil {
//		http.Error(w, "product not found", http.StatusInternalServerError)
//		return
//	}
//}
//
//func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		prod := &data.Product{}
//
//		err := prod.FromJSON(r.Body)
//		if err != nil {
//			p.l.Println("[ERROR] deserializing product", err)
//			http.Error(w, "Error reading product", http.StatusBadRequest)
//			return
//		}
//
//		err = prod.Validate()
//		if err != nil {
//			p.l.Println("[ERROR] validating product", err)
//			http.Error(w, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
//			return
//		}
//
//		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
//		r = r.WithContext(ctx)
//
//		next.ServeHTTP(w, r)
//	})
//}
