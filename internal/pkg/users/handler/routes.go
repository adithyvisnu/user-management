package handler

import "net/http"

/*
RegisterHandler function is for registering any handler inside user package
*/
func RegisterHandler(server *http.ServeMux) {
	server.HandleFunc("/", Home)
	server.HandleFunc("/api/login/v1", Login)
}
