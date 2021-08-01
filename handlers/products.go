package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Faizan-Zia/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
	} else if r.Method == http.MethodPost {
		p.AddProduct(rw, r)
	} else if r.Method == http.MethodPut {
		rx := regexp.MustCompile("/([0-9]+)")
		g := rx.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid Request URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid Request URI", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(g[0][1])
		if err != nil {
			http.Error(rw, "Invalid Request URI", http.StatusBadRequest)
		}
		p.l.Println("Got id: ", id)
		p.UpdateProduct(rw, r, id)

	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Requests")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
	} else {
		data.AddProduct(prod)
		p.l.Printf("Product: %#v", prod)
		rw.WriteHeader(http.StatusCreated)

	}
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request, id int) {
	p.l.Println("Handle PUT Requests")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
	} else {
		err := data.UpdateProduct(prod, id)
		if err == data.ErrProductNotFound {
			http.Error(rw, "Product Not Found", http.StatusNotFound)
			return
		}
	}
}
