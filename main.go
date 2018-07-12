package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// j, _ := json.Marshal(d)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"text": "helloworld"}`))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", HelloWorldHandler).Methods("GET")

	if err := http.ListenAndServe(":3030", r); err != nil {
		log.Fatal(err)
	}
}
