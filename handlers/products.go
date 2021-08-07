package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Faizan-Zia/microservices/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Requests")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Requests")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
	p.l.Printf("Product: %#v", prod)
	rw.WriteHeader(http.StatusCreated)

}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Requests")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	err = data.UpdateProduct(prod, id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

}

type KeyProduct struct{}

func (p *Products) ProductValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		if err := prod.FromJSON(r.Body); err != nil {
			p.l.Println("[ERROR] Deserializing Product")
			http.Error(
				rw,
				"Unable to decode json",
				http.StatusBadRequest,
			)
			return
		}
		if err := prod.Validate(); err != nil {
			p.l.Println("[ERROR] Deserializing Product")
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %v", err),
				http.StatusBadRequest,
			)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
