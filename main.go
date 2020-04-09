package main

import (
	"fmt"
	"net/http"
	"os"

	"levpay/controllers"

	"github.com/gorilla/mux"
)

func main() {

	//dbName := os.Getenv("db_name")
	//fmt.Println(dbName)
	router := mux.NewRouter()
	router.HandleFunc("/api/heroes", controllers.GetHeroes).Methods("GET")
	router.HandleFunc("/api/heroes/{id}", controllers.GetHeroe).Methods("GET")
	router.HandleFunc("/api/heroes/name/{name}", controllers.GetNameHeroe).Methods("GET")
	router.HandleFunc("/api/heroes", controllers.CreateHeroe).Methods("POST")
	router.HandleFunc("/api/heroes/{id}", controllers.UpdateHeroe).Methods("PUT")
	router.HandleFunc("/api/heroes/{id}", controllers.DeleteHeroe).Methods("DELETE")

	//s.HandleFunc("/venues/{id:[0-9]+}", a.UpdateVenue).Methods("PUT")

	router.HandleFunc("/api/teste", controllers.TesteHttp).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)
	//Launch the app, visit localhost:8000/api/heroes
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
