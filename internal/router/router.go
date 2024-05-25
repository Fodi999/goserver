package router

import (
	
	"github.com/gorilla/mux"
	"goserver/internal/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/file", handlers.FileHandler).Methods("GET")
	r.HandleFunc("/data", handlers.CustomDataHandler).Methods("GET")
	r.HandleFunc("/data-post", handlers.CustomDataPostHandler).Methods("POST")
	r.HandleFunc("/data-put", handlers.UpdateCustomDataHandler).Methods("PUT")
	r.HandleFunc("/data-delete", handlers.DeleteCustomDataHandler).Methods("DELETE")
	r.HandleFunc("/data-all", handlers.GetAllCustomDataHandler).Methods("GET")
	r.HandleFunc("/custom-file", handlers.CustomFileHandler).Methods("GET", "POST")
	return r
}

