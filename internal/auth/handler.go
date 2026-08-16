package auth

import (
	"net/http"

	"github.com/Asker231/Velbex/configs"
)


type Auth struct{
	config *configs.Config
}

func InitAuth(app *http.ServeMux, config *configs.Config){
	authHandler := &Auth{
		config: config,
	}

	app.HandleFunc("POST /auth/login",authHandler.login())
	app.HandleFunc("POST /auth/register",authHandler.register())

}


func(a *Auth) login()http.HandlerFunc{
	return  func(w http.ResponseWriter, r *http.Request) {

	}
}
func(a *Auth) register()http.HandlerFunc{
	return  func(w http.ResponseWriter, r *http.Request) {

	}
}