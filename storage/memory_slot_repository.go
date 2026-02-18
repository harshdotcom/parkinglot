package storage

import (
	"errors"
	"strconv"
	"sync"

	"harshdotcom/parkinglot/models"
)

type MemorySlotRepository struct {
	slots []models.Slot
	mu    sync.Mutex
}

func NewMemorySlotRepository() *MemorySlotRepository {
	var slots []models.Slot

	for i := 1; i <= 4; i++ {
		slots = append(slots, models.Slot{
			ID:       i,
			SlotName: "Slot-" + strconv.Itoa(i),
			Distance: i * 10,
			IsBooked: false,
		})
	}

	return &MemorySlotRepository{
		slots: slots,
	}
}

func (m *MemorySlotRepository) GetNearestAvailable() (*models.Slot, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var nearest *models.Slot

	for i := range m.slots {
		slot := &m.slots[i]

		if slot.IsBooked {
			continue
		}

		if nearest == nil || slot.Distance < nearest.Distance {
			nearest = slot
		}
	}

	if nearest == nil {
		return nil, errors.New("no slots available")
	}

	nearest.IsBooked = true
	return nearest, nil
}
