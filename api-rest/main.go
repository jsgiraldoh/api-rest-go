package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	server := http.ListenAndServe(":8080", router)
	//Obtener los logs del servidor
	log.Fatal(server)
}
