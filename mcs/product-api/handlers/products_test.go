package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Raghav847/go-by-example-practice/mcs/product-api/data"
)

func testProductMux() *http.ServeMux {
	ph := NewProducts(log.New(io.Discard, "", 0))
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", ph.GetProducts)
	mux.Handle("POST /", ph.MiddlewareValidateProduct(http.HandlerFunc(ph.AddProducts)))
	mux.Handle("PUT /{id}", ph.MiddlewareValidateProduct(http.HandlerFunc(ph.UpdateProducts)))
	return mux
}

func TestProductRoutes(t *testing.T) {
	mux := testProductMux()

	t.Run("get products", func(t *testing.T) {
		response := httptest.NewRecorder()
		mux.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/", nil))

		if response.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
		}

		var products data.Products
		if err := json.NewDecoder(response.Body).Decode(&products); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if len(products) == 0 {
			t.Fatal("expected at least one product")
		}
	})

	t.Run("add valid product", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(
			`{"name":"Mocha","price":3.25,"sku":"moc-hot-cup"}`,
		))
		mux.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d: %s", http.StatusOK, response.Code, response.Body.String())
		}

		products := data.GetProducts()
		added := products[len(products)-1]
		if added.Name != "Mocha" || added.SKU != "moc-hot-cup" {
			t.Fatalf("product was not added correctly: %+v", added)
		}
	})

	t.Run("update valid product", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPut, "/1", strings.NewReader(
			`{"name":"Iced Latte","price":3.50,"sku":"ice-lat-cup"}`,
		))
		mux.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d: %s", http.StatusOK, response.Code, response.Body.String())
		}

		updated := data.GetProducts()[0]
		if updated.Name != "Iced Latte" || updated.SKU != "ice-lat-cup" {
			t.Fatalf("product was not updated correctly: %+v", updated)
		}
	})

	t.Run("update missing product", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPut, "/999", strings.NewReader(
			`{"name":"Mocha","price":3.25,"sku":"moc-hot-cup"}`,
		))
		mux.ServeHTTP(response, request)

		if response.Code != http.StatusNotFound {
			t.Fatalf("expected status %d, got %d", http.StatusNotFound, response.Code)
		}
	})

	t.Run("update invalid ID", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPut, "/invalid", strings.NewReader(
			`{"name":"Mocha","price":3.25,"sku":"moc-hot-cup"}`,
		))
		mux.ServeHTTP(response, request)

		if response.Code != http.StatusBadRequest {
			t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
		}
	})
}

func TestProductValidationMiddleware(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{name: "missing name", body: `{"price":3.25,"sku":"moc-hot-cup"}`},
		{name: "invalid price", body: `{"name":"Mocha","price":0,"sku":"moc-hot-cup"}`},
		{name: "invalid SKU", body: `{"name":"Mocha","price":3.25,"sku":"invalid"}`},
		{name: "malformed JSON", body: `{"name":`},
	}

	mux := testProductMux()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.body))
			mux.ServeHTTP(response, request)

			if response.Code != http.StatusBadRequest {
				t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
			}
		})
	}
}
