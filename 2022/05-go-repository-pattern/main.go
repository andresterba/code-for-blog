package main

import (
	"fmt"

	"github.com/andresterba/2022/05-go-repository-pattern/interfaces"
)

func main() {
	userRepository := NewGenericRepository[interfaces.User]()
	petRepository := NewGenericRepository[interfaces.Pet]()

	user1 := NewUser("Peter", []string{"1", "2"})
	petOfPeter := NewPet("Peter's Cat", user1.ID())

	err := userRepository.Add(user1)
	if err != nil {
		panic(err)
	}

	err = petRepository.Add(petOfPeter)
	if err != nil {
		panic(err)
	}

	userFromRepo, err := userRepository.Get(user1.ID())
	if err != nil {
		panic(err)
	}

	petFromRepo, err := petRepository.Get(petOfPeter.ID())
	if err != nil {
		panic(err)
	}

	fmt.Println(userFromRepo)
	fmt.Println(petFromRepo)
}
