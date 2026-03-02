package handlers

import (
	"Warehouse/dto"
	"Warehouse/models"
	"Warehouse/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	warehouse repository.Warehouse
}

func NewHTTPHandlers() HTTPHandlers {
	return HTTPHandlers{
		warehouse: repository.NewWarehouse(),
	}
}

func SendErrorResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	errorResponse := dto.NewErrorResponse(err.Error())
	res, convertErr := json.Marshal(errorResponse)
	if convertErr != nil {
		panic(convertErr)
	}
	if _, writeErr := w.Write(res); writeErr != nil {
		println("Ошибка при записи ответа с информацией об ошибке")
	}
}

func (h *HTTPHandlers) HandleCreateItem(w http.ResponseWriter, r *http.Request) {
	item := models.NewItem()

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := h.warehouse.AddItem(item); err != nil {
		SendErrorResponse(w, err, http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)

	b, convertErr := json.Marshal(item)
	if convertErr != nil {
		panic(convertErr)
	}
	if _, err := w.Write(b); err != nil {
		println("Ошибка при записи ответа с созданным товаром")
	}
}

func (h *HTTPHandlers) HandleDeleteItem(w http.ResponseWriter, r *http.Request) {
	itemId := mux.Vars(r)["item_id"]

	if err := h.warehouse.DeleteItem(itemId); err != nil {
		SendErrorResponse(w, err, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPHandlers) HandleGetItem(w http.ResponseWriter, r *http.Request) {
	itemId := mux.Vars(r)["item_id"]

	item, getErr := h.warehouse.GetItem(itemId)
	if getErr != nil {
		SendErrorResponse(w, getErr, http.StatusNotFound)
		return
	}

	b, convertErr := json.Marshal(item)
	if convertErr != nil {
		panic(convertErr)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		println("Ошибка при записи ответа с полученным товаром")
	}

}
