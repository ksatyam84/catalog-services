package main

import (
	"log"
	"net/http"
	"simple-rest/pkg/controllers"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"simple-rest/pkg/controllers"
)
//func HandleRequest(w )
func main() {
	r := mux.NewRouter().StrictSlash(true)
	//routes.RegisterBookStoreRoutes(r)
	//http.Handle("/", r)
	r.HandleFunc("/product/", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/product/", controllers.GetProduct).Methods("GET")
	r.HandleFunc("/product/{productId}", controllers.GetProductById).Methods("GET")
	r.HandleFunc("/product/{productId}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/{productId}", controllers.DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}