package service

import (
	"harshdotcom/parkinglot/models"
	"harshdotcom/parkinglot/storage"
)

type ParkingService struct {
	repo storage.SlotRepository
}

func NewParkingService(repo storage.SlotRepository) *ParkingService {
	return &ParkingService{
		repo: repo,
	}
}

func (p *ParkingService) AssignSlot() (*models.Slot, error) {
	return p.repo.GetNearestAvailable()
}
