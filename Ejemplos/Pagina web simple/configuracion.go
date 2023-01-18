package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	fmt.Printf("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
