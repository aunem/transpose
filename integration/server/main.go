package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	fmt.Println("starting server...")
	fmt.Println("start of handle")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":80", r))
}

// HomeHandler is the home handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("vars: ", vars)
	fmt.Printf("req: %+v \n", r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home!")
}
