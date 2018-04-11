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
	if err := http.ListenAndServe(":3000", r); err != nil {
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
	response, _ := json.Marshal(books)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
