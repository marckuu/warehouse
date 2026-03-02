package server

import (
	"Warehouse/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerManager struct {
	httpHandlers handlers.HTTPHandlers
}

func NewServerManager() ServerManager {
	return ServerManager{
		httpHandlers: handlers.NewHTTPHandlers(),
	}
}

func (s *ServerManager) StartServer() {
	router := mux.NewRouter()

	router.Path("/items").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateItem)
	router.Path("/items/{item_id}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteItem)
	router.Path("/items").Methods("GET").Queries("weight", "{weight}").HandlerFunc(s.httpHandlers.HandleGetItemsLighterThan)
	router.Path("/items/{item_id}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetItem)
	router.Path("/items").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetAllItems)
	router.Path("/items/{item_id}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleChangeItemTitle)

	err := http.ListenAndServe(":9011", router)
	if err != nil {
		println("Ошибка при запуске сервера")
		return
	}
}
