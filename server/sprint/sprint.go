package command

import "github.com/satori/go.uuid"

type AddSprint struct {
	ID   uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name string    `json:"name" example:"sprint name"`
}
