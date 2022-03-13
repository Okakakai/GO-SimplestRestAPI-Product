package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ItemParams struct {
	Id       string `json:"id"`
	JanCode  string `json:"jan_code,omitempty"`
	ItemName string `json:"item_name,omitempty"`
	Price    int    `json:"price,omitempty"`
}

var items []*ItemParams

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the GO API server")
	fmt.Println("Root endpoint is hooked!")
}

func StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", rootPage)
	router.HandleFunc("/items", fetchAllItems).Methods("GET")
	router.HandleFunc("/item/{id}", fetchSingleItem).Methods("GET")

	router.HandleFunc("/item", createItem).Methods("POST")
	router.HandleFunc("item/{id}", deleteItem).Methods("DELETE")
	router.HandleFunc("item/{id}", updateItem).Methods("PUT")

	return http.ListenAndServe(fmt.Sprint(":%d", 8080), router)

}
