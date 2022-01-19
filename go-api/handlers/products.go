package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/bhongy/kimidori/go-api/product-api/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) Get(rw http.ResponseWriter, r *http.Request) {
	ps := data.GetProducts()
	if err := json.NewEncoder(rw).Encode(ps); err != nil {
		http.Error(rw, "Unable to encode JSON", http.StatusInternalServerError)
	}
}

func (p *Products) Post(rw http.ResponseWriter, r *http.Request) {
	var prod data.Product
	if err := json.NewDecoder(r.Body).Decode(&prod); err != nil {
		http.Error(rw, "Unable to decode request body", http.StatusBadRequest)
		return
	}
	data.AddProduct(&prod)
}

func (p *Products) Put(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Printf("Cannot parse ID path=%v: %v\n", r.URL.Path, err)
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	var prod data.Product
	if err := json.NewDecoder(r.Body).Decode(&prod); err != nil {
		http.Error(rw, "Unable to decode request body", http.StatusBadRequest)
		return
	}

	prod.ID = id
	if err = data.UpdateProduct(&prod); err != nil {
		if errors.Is(err, data.ErrProductNotFound) {
			http.Error(rw, "Not found", http.StatusNotFound)
		} else {
			p.l.Printf("Cannot update product: %v\n", err)
			http.Error(rw, "Unknown error", http.StatusInternalServerError)
		}
	}
}
