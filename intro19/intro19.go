// Package intro19 provides the implementation of the introduction to golang
// (version 2019) example code
package intro19

import (
	"github.com/gorilla/mux"
	"net/http"
)

// RunServer executes the server logic
func RunServer(addr string) error {
	router := mux.NewRouter()

	router.HandleFunc("/hello", helloHandler)
	return http.ListenAndServe(addr, router)
}
