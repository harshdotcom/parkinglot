package main

import (
	"harshdotcom/parkinglot/handlers"
	"net/http"
)

const (
	port = ":9090"
)

func main() {
	http.HandleFunc("/park", handlers.ParkVehicle)
	http.HandleFunc("/leave", handlers.LeaveVehicle)
	http.HandleFunc("/status", handlers.GetStatus)
	http.HandleFunc("/users", handlers.CreateUser)

	http.ListenAndServe(port, nil)
}
