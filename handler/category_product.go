package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sinulingga23/go-pos/definition"
	"github.com/sinulingga23/go-pos/payload"
)

var (
	messageSuccessCreateCategoryProduct string = "Success to create a category product."
)


func (h handler) CreateCategoryProduct(w http.ResponseWriter, r *http.Request) {
	createCategoryProductRequest := payload.CreateCategoryProductRequest{}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		response := struct {
			Message string `json:"message"`
		}{Message: definition.ErrBadRequest.Error()}

		bytes, _ := json.Marshal(response)

		
		w.WriteHeader(http.StatusBadRequest)
		w.Write(bytes)
		return
	}

	if err := json.Unmarshal(bytes, &createCategoryProductRequest); err != nil {
		log.Printf("[ERROR]: %v\n", err)
		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}

		bytes, _ := json.Marshal(response)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(bytes)
		return
	}

	createdCategoryProduct, err := h.categoryProductService.Create(createCategoryProductRequest)
	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		if err == definition.ErrBadRequest {
			response := struct {
				Message string `json:"message"`
			}{Message: err.Error()}
	
			bytes, _ := json.Marshal(response)

			w.WriteHeader(http.StatusBadRequest)
			w.Write(bytes)
			return
		}

		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}

		bytes, _ := json.Marshal(response)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bytes)
		return
	}

	response  := struct {
		Message string `json:"message"`
		Data payload.CategoryProduct `json:"data"`
	}{Message: "Success to create a category product.", Data: *createdCategoryProduct}

	bytes, _ = json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}
