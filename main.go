package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",	// TODO: load from .env and check dev/prod mode
		Handler: http.NewServeMux(),
	}

	fmt.Println("Starting server")
	server.ListenAndServe()
}
