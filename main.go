package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	m := mux.NewRouter()
	server := &http.Server{
		Handler:     m,
		Addr: 		"localhost"+":8080",
	}

	m.HandleFunc("/", Handle)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}


}

func Handle(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from server"))
}