package handlers

import (
	"Warehouse/dto"
	"Warehouse/models"
	"Warehouse/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	if convertErr := json.NewEncoder(w).Encode(errorResponse); convertErr != nil {
		fmt.Println("Ошибка при записи ответа с информацией об ошибке")
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

	if convertErr := json.NewEncoder(w).Encode(item); convertErr != nil {
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

	w.WriteHeader(http.StatusOK)
	if convertErr := json.NewEncoder(w).Encode(item); convertErr != nil {
		println("Ошибка при записи ответа с полученным товаром")
	}
}

func (h *HTTPHandlers) HandleGetAllItems(w http.ResponseWriter, r *http.Request) {
	items := h.warehouse.GetAllTItems()

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		println("Ошибка при записи ответа со всеми товарами")
	}
}

func (h *HTTPHandlers) HandleGetItemsLighterThan(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	weight := query.Get("weight")

	weightConverted, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		SendErrorResponse(w, err, http.StatusInternalServerError)
	}

	filteredItems := h.warehouse.GetItemLighterThan(weightConverted)

	w.WriteHeader(http.StatusOK)
	if convertErr := json.NewEncoder(w).Encode(filteredItems); convertErr != nil {
		fmt.Println("Ошибка при записи ответа с товарами легче указанного веса")
	}
}

func (h *HTTPHandlers) HandleChangeItemTitle(w http.ResponseWriter, r *http.Request) {
	itemId := mux.Vars(r)["item_id"]

	request := dto.NewChangeItemTitleRequest("")
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		SendErrorResponse(w, err, http.StatusBadRequest)
	}

	item, err := h.warehouse.ChangeItemTitle(itemId, request.Title)
	if err != nil {
		SendErrorResponse(w, err, http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	if convertErr := json.NewEncoder(w).Encode(item); convertErr != nil {
		fmt.Println("Ошибка при записи ответа с измененым товаром")
	}
}
