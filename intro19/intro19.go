// Package intro19 provides the implementation of the introduction to golang
// (version 2019) example code
package intro19

import "net/http"

// RunServer executes the server logic
func RunServer(addr string) error {
	return http.ListenAndServe(addr, nil)
}
