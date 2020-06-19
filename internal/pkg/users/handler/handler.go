package handler

import (
	"net/http"

	"github.com/adithyvisnu/user-management/internal/app/delivery"
)

/*
Home is for showing for welcome page
*/
func Home(writer http.ResponseWriter, req *http.Request) {
	delivery.JSONResponse(writer, []byte("Hello"))
}

/*
Login handles JWT for login method OAuth2
*/
func Login(writer http.ResponseWriter, req *http.Request) {
	delivery.JSONResponse(writer, "Login API")
}
