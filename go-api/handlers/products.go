package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"

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
		p.l.Println("Handle GET Products")
		p.get(rw, r)
	case http.MethodPost:
		p.l.Println("Handle POST Products")
		p.post(rw, r)
	case http.MethodPut:
		p.l.Println("Handle PUT Products")
		p.put(rw, r)
	default:
		status := http.StatusMethodNotAllowed
		http.Error(rw, http.StatusText(status), status)
	}
}

func (p *Products) get(rw http.ResponseWriter, r *http.Request) {
	ps := data.GetProducts()
	if err := json.NewEncoder(rw).Encode(ps); err != nil {
		http.Error(rw, "Unable to encode JSON", http.StatusInternalServerError)
	}
}

func (p *Products) post(rw http.ResponseWriter, r *http.Request) {
	var prod data.Product
	if err := json.NewDecoder(r.Body).Decode(&prod); err != nil {
		http.Error(rw, "Unable to decode request body", http.StatusBadRequest)
		return
	}
	data.AddProduct(&prod)
}

func (p *Products) put(rw http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path)
	if err != nil {
		p.l.Printf("Cannot parse ID path=%v: %v\n", r.URL.Path, err)
		http.Error(rw, "Invalid URI", http.StatusBadRequest)
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

func parseID(path string) (int, error) {
	re := regexp.MustCompile("/([0-9]+)")
	g := re.FindAllStringSubmatch(path, -1)
	if len(g) != 1 {
		return 0, errors.New("Matches multiple")
	}

	id, err := strconv.Atoi(g[0][1])
	if err != nil {
		return 0, err
	}

	return id, nil
}
