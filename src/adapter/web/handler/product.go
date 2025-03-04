package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/adapter/dto"
	"github.com/silverioTenor/fc-arq-hexagonal-go/src/app"
)

func MakeProductHandlers(r *mux.Router, n * negroni.Negroni, service app.IProductService) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(get(service)),
	)).Methods("GET", "OPTIONS")
	
	r.Handle("/product", n.With(
		negroni.Wrap(create(service)),
	)).Methods("POST", "OPTIONS")
	
	r.Handle("/product/{id}/change-status", n.With(
		negroni.Wrap(changeStatus(service)),
	)).Methods("PATCH", "OPTIONS")
}

func get(service app.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func create(service app.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDto dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func changeStatus(service app.IProductService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var productDto dto.Product
		err = json.NewDecoder(r.Body).Decode(&productDto)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		newProduct := app.Product{
			Id:     product.GetId(),
			Name:   product.GetName(),
			Price:  0,
			Status: productDto.Status,
		}

		product, err = service.Toggle(&newProduct)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}