package topic

import "github.com/satori/go.uuid"

type topicModel struct {
	ID     uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Sprint uuid.UUID `json:"sprint" example:"550e8400-e29b-41d4-a716-446655440001" format:"uuid"`
	Text   string    `json:"text" example:"a short topic description"`
}
