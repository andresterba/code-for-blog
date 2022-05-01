package main

import (
	"fmt"

	"github.com/andresterba/2022/05-go-repository-pattern/interfaces"
	"github.com/google/uuid"
)

type Pet struct {
	id    uuid.UUID
	Name  string
	Owner uuid.UUID
}

func NewPet(name string, owner uuid.UUID) interfaces.Pet {
	return &Pet{
		id:    uuid.New(),
		Name:  name,
		Owner: owner,
	}
}

func (p *Pet) ID() uuid.UUID {
	return p.id
}

func (p *Pet) String() string {
	return fmt.Sprintf("I'm %s and I belong to %s", p.Name, p.Owner)
}
