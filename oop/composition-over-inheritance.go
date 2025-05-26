package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "Some sound"
}

type Dog struct {
	Animal
	Breed string
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Zen"},
		Breed:  "Golden Retriever",
	}

	fmt.Println(d.Name)
	fmt.Println(d.Speak())
}
