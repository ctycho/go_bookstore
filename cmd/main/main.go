package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/ctycho/go_bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	fmt.Println("Listening server on PORT:9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", nil))
}