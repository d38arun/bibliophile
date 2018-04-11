package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	. "../config"
	. "../mapper"
	"log"
	"github.com/gin-gonic/gin/json"
)

func Route() {
	r := mux.NewRouter()
	r.HandleFunc("/",DisplayBooks).Methods("GET")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}

var config = Config{}
var dao = BookDAO{}

func DisplayBooks(writer http.ResponseWriter, request *http.Request) {
	books, err := dao.FindAllBooks()
	if err != nil {
		log.Fatal(err)
	}
	respondWithJson(writer, books)
}

func respondWithJson(writer http.ResponseWriter, books interface{}) {
	response, _ := json.Marshal(books)
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Cache-Control, Pragma, Origin, Authorization, Content-Type, X-Requested-With")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
