package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Office struct {
	Id      string   `json:"id"`
	Isbn    string   `json:"isbn"`
	Branch  string   `jason:"branch"`
	Manager *Manager `jason:"director"`
}

type Manager struct {
	Firstname string `jason:"firstname"`
	Lastname  string `jason:"lastname"`
}

var offices []Office

func getOffices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offices)

}

func deleteOffice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	temp := mux.Vars(r)
	for index, item := range offices {
		if item.Id == temp["id"] {
			offices = append(offices[:index], offices[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(offices)
}

func getOffice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	temp := mux.Vars(r)
	for _, item := range offices {
		if item.Id == temp["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createOffice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var office Office
	_ = json.NewDecoder(r.Body).Decode(&office)
	office.Id = strconv.Itoa(rand.Intn(10))
	offices = append(offices, office)
	json.NewEncoder(w).Encode(office)
}

func updateOffice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	temp := mux.Vars(r)
	for index, item := range offices {
		if item.Id == temp["id"] {
			offices = append(offices[:index], offices[index+1:]...)
			var office Office
			_ = json.NewDecoder(r.Body).Decode(&office)
			office.Id = temp["id"]
			offices = append(offices, office)
			json.NewEncoder(w).Encode(office)
		}
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/offices", getOffices).Methods("GET")
	r.HandleFunc("/offices/{id}", getOffice).Methods("GET")
	r.HandleFunc("/offices", createOffice).Methods("POST")
	r.HandleFunc("/offices/{id}", updateOffice).Methods("PUT")
	r.HandleFunc("/offices/{id}", deleteOffice).Methods("DELETE")

	fmt.Println("Starting Server at Port 3306")
	log.Fatal(http.ListenAndServe(":3306", r))
}
