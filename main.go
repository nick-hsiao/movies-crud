package main

import (
	_ "encoding/json"
	"fmt"
	_ "log"
	_ "math/rand"
	"net/http"
	_ "strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
}

func main() {

	router := mux.NewRouter()

	//specify endpoints, handler functions and HTTP method

	//router.HandleFunc("/movie", Movie).Methods("GET")
	http.Handle("/", router)
	fmt.Print("working")
	//start and listen to requests
	http.ListenAndServe(":8080", router)

}
