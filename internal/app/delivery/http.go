package delivery

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status int
	Data   interface{}
}

/*
JSONResponse function is used for http response which
shows the browser for result in JSON format
*/
func JSONResponse(writer http.ResponseWriter, data interface{}) {
	res := response{}
	res.Data = data
	res.Status = http.StatusOK
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(res)
}

/*
XMLResponse function sends an XML type for caller
*/
func XMLResponse(writer http.ResponseWriter, data interface{}) {

}

/*
BlobResponse function sends an blob file which downloadable
*/
func BlobResponse(writer http.ResponseWriter, data interface{}) {

}
