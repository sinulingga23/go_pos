package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sinulingga23/go-pos/definition"
	"github.com/sinulingga23/go-pos/payload"
)


var (
	messageSuccessCreateProduct string = "Success to create a product."
	messageAddImagesInProcess string = "Add Images in Process."
)

func (h handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		log.Printf("[ProductHandler][r.ParseForm]: %s", err.Error())

		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}
		
		bytes, _ := json.Marshal(response)
		w.Write(bytes)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(r.ContentLength); err != nil {
		log.Printf("[ProductHandler][r.ParseMultipartForm]: %s", err.Error())

		respnose := struct {
			Message string `json:"message"`
		}{Message: err.Error()}

		bytes, _ := json.Marshal(respnose)
		w.Write(bytes)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stock, err := strconv.Atoi(r.PostFormValue("stock"))
	if err != nil {
		log.Printf("[ProductHandler][strconv.Atoi]: %s", err.Error())

		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}

		bytes, _ := json.Marshal(response)
		w.Write(bytes)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		log.Printf("[ProductHandler][strconv.ParseFloat]: %s", err.Error())

		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}

		bytes, _ := json.Marshal(response)
		w.Write(bytes)
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


		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}
		bytes, _ := json.Marshal(response)

		if err == definition.ErrBadRequest {
			w.Write(bytes)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err == definition.ErrDataNotFound {
			w.Write(bytes)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Write(bytes)
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

func (h handler) AddImagesToProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if err := r.ParseMultipartForm(r.ContentLength); err != nil {
		log.Printf("[ProductHandler][r.ParseMultipartForm]: %s", err.Error())
		
		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}

		bytes, _ := json.Marshal(response)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(bytes)
		return
	}

	 productId := chi.URLParam(r, "id")
	 if err := h.productService.AddImagesToProduct(productId, payload.AddImageToProductRequest{
		 ProudctId: r.PostFormValue("productId"),
		 Images: r.MultipartForm.File,
	 }); err != nil {
		 log.Printf("[ProductHandler][h.productService.AddImagesToProduct]: %s", err.Error())

		 response := struct {
			 Message string `json:"message"`
		 }{Message: err.Error()}
		 bytes, _ := json.Marshal(response)

		 if err == definition.ErrBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(bytes) 
			return
		 }

		 w.WriteHeader(http.StatusInternalServerError)
		 w.Write(bytes)
		 return
	 }

	 response := struct {
		 Message string `json:"message"`
	 }{Message: messageAddImagesInProcess}
	 bytes, _ := json.Marshal(response)
	 w.WriteHeader(http.StatusOK)
	 w.Write(bytes)
	 return
}