package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	//"html"
	//"github.com/joho/godotenv"
	//"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:ezaz@1234@tcp(host.docker.internal:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local"

type Office struct {
	Id     string `json:"id"`
	Isbn   string `json:"isbn"`
	Branch string `jason:"branch"`
	//Manager *Manager `jason:"director"`
}

// type Manager struct {
// 	Firstname string `jason:"firstname"`
// 	Lastname  string `jason:"lastname"`
// }

var offices []Office

func getOffices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DB.Find(&offices)
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
		DB.Delete(&offices, temp["id"])
	}
	json.NewEncoder(w).Encode(offices)
}

func getOffice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	temp := mux.Vars(r)
	for _, item := range offices {
		if item.Id == temp["id"] {
			DB.First(&item, temp["id"])
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
	DB.Create(&office)
	json.NewEncoder(w).Encode(office)
}

func updateOffice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	temp := mux.Vars(r)
	for index, item := range offices {
		if item.Id == temp["id"] {
			offices = append(offices[:index], offices[index+1:]...)
			var office Office
			DB.First(&office, temp["id"])
			_ = json.NewDecoder(r.Body).Decode(&office)
			office.Id = temp["id"]
			offices = append(offices, office)
			DB.Save(&office)
			json.NewEncoder(w).Encode(office)
		}
	}
}

func DataConnect() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Office{
		Id:     "",
		Isbn:   "",
		Branch: "",
		//Manager: &Manager{},
	})
}

func middleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("Islam ")
		handler.ServeHTTP(response, request)
	}
}

func middlewaredo(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("Ezazul")
		handler.ServeHTTP(response, request)
	}
}
func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprintf(w, "Hi")
	// })

	DataConnect()

	r := mux.NewRouter()

	r.HandleFunc("/offices", middlewaredo(middleware(getOffices))).Methods("GET")
	r.HandleFunc("/offices/{id}", middlewaredo(middleware(getOffice))).Methods("GET")
	r.HandleFunc("/offices", middlewaredo(middleware(createOffice))).Methods("POST")
	r.HandleFunc("/offices/{id}", middlewaredo(middleware(updateOffice))).Methods("PUT")
	r.HandleFunc("/offices/{id}", middlewaredo(middleware(deleteOffice))).Methods("DELETE")

	fmt.Println("Starting Server at Port 8080")
	log.Fatal(http.ListenAndServe(":8081", r))
}
