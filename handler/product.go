package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sinulingga23/go-pos/definition"
	"github.com/sinulingga23/go-pos/payload"
)


var (
	messageSuccessCreateProduct string = "Success to create a product."
)

func (h handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		log.Printf("[ProductHandler][r.ParseForm]: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(r.ContentLength); err != nil {
		log.Printf("[ProductHandler][r.ParseMultipartForm]: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stock, err := strconv.Atoi(r.PostFormValue("stock"))
	if err != nil {
		log.Printf("[ProductHandler][strconv.Atoi]: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		log.Printf("[ProductHandler][strconv.ParseFloat]: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	
	createdProduct, err := h.productService.Create(payload.CreateProductRequest{
		CategoryProductIds: r.PostForm["categoryProductsIds[]"],
		ProductName: r.PostFormValue("productName"),
		Description: r.PostFormValue("description"),
		Stock: stock,
		Price: price,
		Images: r.MultipartForm.File,
	})
	if err != nil {
		log.Printf("[ProductHandler][h.productService.Create]: %s", err.Error())

		if err == definition.ErrBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err == definition.ErrDataNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
		Data payload.Product `json:"data"`
	}{Message: messageSuccessCreateProduct, Data: *createdProduct}
	bytes, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}