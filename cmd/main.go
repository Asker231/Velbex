package main

import (
	"net/http"

	"github.com/Asker231/Velbex/configs"
)

func main() {
	config := configs.InitConfig()
	
	app := http.NewServeMux()
	
	server := http.Server{
		Addr: ":8081",
		Handler: app,
	}

	server.ListenAndServe()
}