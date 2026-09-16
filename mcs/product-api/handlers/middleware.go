package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Raghav847/go-by-example-practice/mcs/product-api/data"
)

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "error deserializing product", http.StatusBadRequest)
			return
		}

		errs := p.v.Validate(prod)
		if len(errs) > 0 {
			p.l.Println("[ERROR] validating product", err)
			http.Error(w, fmt.Sprintf("error validating product: %v", errs.Errors()), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
