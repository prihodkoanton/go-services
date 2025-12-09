package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("User service started on port 9090")

	http.HandleFunc("/health", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":9090", nil)
}
