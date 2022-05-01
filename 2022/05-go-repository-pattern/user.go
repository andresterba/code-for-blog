package main

import (
	"fmt"

	"github.com/andresterba/2022/05-go-repository-pattern/interfaces"
	"github.com/google/uuid"
)

type User struct {
	id      uuid.UUID
	Name    string
	Hobbies []string
}

func NewUser(name string, hobbies []string) interfaces.User {
	return &User{
		id:      uuid.New(),
		Name:    name,
		Hobbies: hobbies,
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) String() string {
	return fmt.Sprintf("I'm %s and I like %s", u.Name, u.Hobbies)
}
