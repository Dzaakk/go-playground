package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}

func (p *Person) HaveBirthday() {
	p.Age++
}
