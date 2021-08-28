package go_designpattern

import "fmt"

type AnimalType int

const (
	AnimalTypeCat = 0
	AnimalTypeDog = 1
)

type Animal interface {
	SayHello() string
}

type Dog struct {
}

func (dog *Dog) SayHello() string {
	return fmt.Sprintf("hello, I'm a %s", "dog")
}

type Cat struct {
}

func (cat *Cat) SayHello() string {
	return fmt.Sprintf("hello, I'm a %s", "cat")
}

func NewAnimal(animalType AnimalType) (Animal, error) {
	switch animalType {
	case AnimalTypeCat:
		return &Cat{}, nil
	case AnimalTypeDog:
		return &Dog{}, nil
	default:
		return nil, fmt.Errorf("unknown animalType: %v", animalType)
	}
}
