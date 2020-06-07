package main

import (
	"fmt"
	"log"
	"net/http"

	user "github.com/adithyvisnu/user-management/internal/pkg/users/handler"
)

func main() {
	fmt.Println("---------------------------")
	fmt.Println("> User management started >")
	fmt.Println("---------------------------")

	handler := http.NewServeMux()
	user.RegisterHandler(handler)
	fmt.Println(handler)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
