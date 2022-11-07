package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dragranzer/BMS-MVC/pkg/config"
	"github.com/dragranzer/BMS-MVC/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	config.Connect()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server Run Pada Port : 9010")
	log.Fatal(http.ListenAndServe(":9010", r))
}
