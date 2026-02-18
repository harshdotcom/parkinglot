package storage

import (
	"harshdotcom/parkinglot/models"
)

type SlotRepository interface {
	GetNearestAvailable() (*models.Slot, error)
}
