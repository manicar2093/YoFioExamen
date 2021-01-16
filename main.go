package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	http.ListenAndServe(":8000", r)
	fmt.Println("Main! :D")
}

func init() {
	fmt.Println("Init! :D")
}
