package gosandpit

import (
	"fmt"
)

// Animal ...
type Animal interface {
	Speak() string
}

// Dog ...
type Dog struct {
}

// Speak dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat struct
type Cat struct {
}

// Speak cat
func (c Cat) Speak() string {
	return "Meow!"
}

// Llama struct
type Llama struct {
}

// Speak Llama
func (l Llama) Speak() string {
	return "?????"
}

// JavaProgrammer ...
type JavaProgrammer struct {
}

// Speak JavaProgrammer
func (j JavaProgrammer) Speak() string {
	return "Design patterns!!"
}

// InterfacesAdv ...
func InterfacesAdv() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
