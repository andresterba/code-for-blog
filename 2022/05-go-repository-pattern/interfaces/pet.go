package interfaces

import "github.com/google/uuid"

type Pet interface {
	ID() uuid.UUID
}
