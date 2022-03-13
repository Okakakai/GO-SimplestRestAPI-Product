package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// json: ~ をパラメータに付与すると、jsonエンコード時にパラメータ名を指定することができます。
// また、omitemptyを付与するとパラメータが空のときに、jsonのパラメータから消すことができます。
// これはクライアントアプリと仕様を統一する必要があります。
type ItemParams struct {
	Id           string    `json:"id"`
	JanCode      string    `json:"jan_code,omitempty"`
	ItemName     string    `json:"item_name,omitempty"`
	Price        int       `json:"price,omitempty"`
	CategoryId   int       `json:"category_id,omitempty"`
	SeriesId     int       `json:"series_id,omitempty"`
	Stock        int       `json:"stock,omitempty"`
	Discontinued bool      `json:"discontinued"`
	ReleaseDate  time.Time `json:"release_date,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
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

func init() {
	items = []*ItemParams{
		&ItemParams{
			Id:           "1",
			JanCode:      "327390283080",
			ItemName:     "item_1",
			Price:        2500,
			CategoryId:   1,
			SeriesId:     1,
			Stock:        100,
			Discontinued: false,
			ReleaseDate:  time.Now(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    time.Now(),
		},
		&ItemParams{
			Id:           "2",
			JanCode:      "3273902878656",
			ItemName:     "item_2",
			Price:        1200,
			CategoryId:   2,
			SeriesId:     2,
			Stock:        200,
			Discontinued: false,
			ReleaseDate:  time.Now(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    time.Now(),
		},
	}
}
