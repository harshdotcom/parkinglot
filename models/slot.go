package models

type Slot struct {
	ID       int    `json:"id"`
	SlotName string `json:"slot_name"`
	Distance int    `json:"distance"`
	IsBooked bool   `json:"is_booked"`
}
