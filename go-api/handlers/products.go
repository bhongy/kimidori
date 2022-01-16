package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bhongy/kimidori/go-api/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.get(rw, r)
	case http.MethodPost:
		// p.post(rw, r)
	default:
		status := http.StatusMethodNotAllowed
		http.Error(rw, http.StatusText(status), status)
	}
}

func (p *Products) get(rw http.ResponseWriter, r *http.Request) {
	ps := data.GetProducts()
	err := json.NewEncoder(rw).Encode(ps)
	if err != nil {
		http.Error(rw, "Unable to encode JSON", http.StatusInternalServerError)
	}
}
