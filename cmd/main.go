package main

import "net/http"

func main() {
	app := http.NewServeMux()
	
	server := http.Server{
		Addr: ":8081",
		Handler: app,
	}

	server.ListenAndServe()
}