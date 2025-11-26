package main

import (
	"fmt"
	"net/http"
)

func main() {
	storage := SetupStorage()
	server := SetupServer(*storage)
	router := http.NewServeMux()

	router.HandleFunc("/", server.HandleTodo)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}
