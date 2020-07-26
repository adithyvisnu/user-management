package delivery

import "net/http"

// JSONResponse gives the response with JSON structures on an API Call
func JSONResponse(writer http.ResponseWriter, data []byte) {
	writer.WriteHeader(200)
	writer.Write(data)
}
