package interfaces

import "github.com/google/uuid"

type User interface {
	ID() uuid.UUID
}


