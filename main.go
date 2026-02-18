package main

import (
	"harshdotcom/parkinglot/handlers"
	"harshdotcom/parkinglot/service"
	"harshdotcom/parkinglot/storage"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	repo := storage.NewMemorySlotRepository()
	service := service.NewParkingService(repo)
	handler := handlers.NewParkingHandler(service)
	http.HandleFunc("/assign", handler.AssignSlot)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
