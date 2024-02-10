package response

import "github.com/google/uuid"

type UserDTO struct {
	ID        uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name      string    `json:"name" example:"Eduardo Melo"`
	Email     string    `json:"email" example:"dudu@gmail.com"`
	Password  string    `json:"password" example:"123456"`
	Admin     bool      `json:"admin" example:"false"`
	CreatedAt string    `json:"created_at" example:"2021-08-01T00:00:00Z"`
	UpdatedAt string    `json:"updated_at" example:"2021-08-01T00:00:00Z"`
}
