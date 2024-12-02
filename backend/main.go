package main

import (
	"fmt"
	"log"
	"net/http"

	"restaurant/database"

	"github.com/gorilla/mux"
)

func main() {
	database.DatabaseConnection()
	fmt.Println("The database connected successfully")

	r := mux.NewRouter()
	r.HandleFunc("/addfood", food.addFood).Methods("UPDATE")
	log.Fatal(http.ListenAndServe(":7000", r))
}
