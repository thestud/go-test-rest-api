package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type crypto struct {
	ID          uint64 `json:ID`
	Name        string `json:"Name"`
	Percentage  uint64 `json:"Percentage"`
}

type allCrypto []crypto

var cryptos = allCrypto{}

func loadJSON() {
	// read file
    data, err := ioutil.ReadFile("./crypto.json")
    if err != nil {
      fmt.Print(err)
    }

    // unmarshall it
    err = json.Unmarshal(data, &cryptos)
    if err != nil {
        fmt.Println("error:", err)
    }
}

func getAllCryptos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cryptos)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to vaporware crypto!")
}

func main() {
	loadJSON()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/cryptos", getAllCryptos).Methods("GET")
	var port = "4000"
	log.Println("Starting Server http://localhost:"+port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}