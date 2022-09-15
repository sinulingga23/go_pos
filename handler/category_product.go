package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sinulingga23/go-pos/definition"
	"github.com/sinulingga23/go-pos/payload"
)

var (
	messageSuccessCreateCategoryProduct string = "Success to create a category product."
	messageSuccesFindCategoryProduct string = "Success to found the category product."
	messageSuccesUpdateCategoryProduct string = "Success to update the category product."
	messageSuccesDeleteCategoryProduct string = "Success to delete the category product."
	messageSuccesGetAllCategoryProduct string = "Success to get all the category product."
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

	response  := struct {
		Message string `json:"message"`
		Data payload.CategoryProduct `json:"data"`
	}{Message: "Success to create a category product.", Data: *createdCategoryProduct}

	bytes, _ = json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}

func (h handler) FindCategoryById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	currentCategoryProduct, err := h.categoryProductService.FindById(id)
	if err != nil {
		log.Printf("[ERROR]: %v\n", err)

		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}
		bytes, _ := json.Marshal(response)

		if err == definition.ErrDataNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write(bytes)
			return
		}

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
		Data payload.CategoryProduct `json:"data"`
	}{Message: messageSuccesFindCategoryProduct, Data: *currentCategoryProduct}

	bytes, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}

func (h handler) UpdateCategoryProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	updateCategoryProductRequest := payload.UpdateCategoryProductRequest{}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR]: %v\n", err)

		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}

		bytes, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(bytes)
		return
	}

	if err := json.Unmarshal(bytes, &updateCategoryProductRequest); err != nil {
		log.Printf("[ERROR]: %v\n", err)

		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}

		bytes, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(bytes)
		return
	}

	updatedCategoryProduct, err := h.categoryProductService.UpdateById(id, updateCategoryProductRequest)
	if err != nil {
		log.Printf("[ERROR]: %v\n", err)

		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}
		bytes, _ := json.Marshal(response) 

		if err == definition.ErrBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(bytes)
			return
		}

		if err == definition.ErrDataNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write(bytes)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bytes)
		return
	}

	response := struct {
		Message string `json:"message"`
		Data payload.CategoryProduct `json:"data"`
	}{Message: messageSuccesUpdateCategoryProduct, Data: *updatedCategoryProduct}

	bytes, _ = json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}

func (h handler) DeleteCategoryProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	id := chi.URLParam(r, "id")

	deleteCategoryProductRequest := payload.DeleteCategoryProductRequest{}
	bytes, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(bytes, &deleteCategoryProductRequest); err != nil {
		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}
		bytes, _ := json.Marshal(response)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(bytes)
		return
	}
	
	if err := h.categoryProductService.DeleteById(id, deleteCategoryProductRequest); err != nil {
		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}
		bytes, _ := json.Marshal(response)

		if err == definition.ErrBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(bytes)
			return
		}

		if err == definition.ErrDataNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write(bytes)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bytes)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{Message: messageSuccesDeleteCategoryProduct}
	bytes, _ = json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}

func (h handler) GetAllCategoryProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	categoryProducts, err := h.categoryProductService.FindAll()
	if err != nil {
		response := struct {
			Message string `json:"message"`
		}{Message: err.Error()}
		bytes, _ := json.Marshal(response)

		if err == definition.ErrDataNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write(bytes)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bytes)
		return
	}

	response := struct {
		Message string `json:"message"`
		Data []*payload.CategoryProduct `json:"data"`
	}{Message: messageSuccesGetAllCategoryProduct, Data: categoryProducts}
	bytes, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}