package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to this life-changing API.\n")
		fmt.Fprintf(w, "Welcome to this life-changing API.\n")
		fmt.Fprintf(w, "Welcome to this life-changing API.\n")
		fmt.Fprintf(w, "Welcome to this life-changing API.\n")
		fmt.Fprintf(w, "Welcome to this life-changing API.\n")
		fmt.Fprintf(w, "Welcome to this life-changing API.\n")
	})

	fmt.Println("Server listening!")

	errs := make(chan error)
	select {
	case errs <- http.ListenAndServe(":8080", r):
		err := <-errs
		fmt.Println("got error from http:", err)
	}
}
