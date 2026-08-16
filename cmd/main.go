package main

import (
	"net/http"

	"github.com/Asker231/Velbex/configs"
	"github.com/Asker231/Velbex/internal/auth"
)

func main() {
	cnf := configs.InitConfig()

	app := http.NewServeMux()
	
	auth.InitAuth(app,cnf)

	server := http.Server{
		Addr: ":8081",
		Handler: app,
	}

	server.ListenAndServe()
}