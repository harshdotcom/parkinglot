package handlers

import (
	"encoding/json"
	"harshdotcom/parkinglot/service"
	"net/http"
)

type ParkingHandler struct {
	service *service.ParkingService
}

func NewParkingHandler(service *service.ParkingService) *ParkingHandler {
	return &ParkingHandler{
		service: service,
	}
}

func (h *ParkingHandler) AssignSlot(w http.ResponseWriter, r *http.Request) {
	slot, err := h.service.AssignSlot()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slot)
}

func ParkVehicle(w http.ResponseWriter, r *http.Request) {

}

func LeaveVehicle(w http.ResponseWriter, r *http.Request) {

}

func GetStatus(w http.ResponseWriter, r *http.Request) {

}
