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
