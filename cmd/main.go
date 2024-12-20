package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	handler "github.com/hwtdspprcmpltl/calcyandex/internal/handler"
)

//curl --location http://localhost:8080 --header "Content-Type: application/json" --data "{\"expression\": \"2+2*2\"}"

func main() {
	http.HandleFunc("/api/v1/calculate", handler.HandleCalculator)

	port := flag.Int("port", 8080, "Port to run the server on")

	flag.Parse()

	log.Printf("привет:) запустил сервер на %d\n", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("не удалось запустить сервер %v", err)
	}
}
