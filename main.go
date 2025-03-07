package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sahika2000/GO_Test-http-server/handlers"
)

func main() {

	// Define routes

	http.HandleFunc("/", handlers.HomeHandler)

	http.HandleFunc("/hello", handlers.HelloHandler)

	http.HandleFunc("/user", handlers.UserHandler)

	port := "8080"

	fmt.Println("Server is running on port", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
