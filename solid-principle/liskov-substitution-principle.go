package main

import "fmt"

type Bird interface {
	Fly() string
}

type Sparrow struct{}

func (s Sparrow) Fly() string {
	return "Sparrow is flying"
}

func LetBirdFly(b Bird) {
	fmt.Println(b.Fly())
}

// ----------------------------
// wrong example
type Ostrich struct{}

func (o Ostrich) Fly() string {
	return "I can't fly"
}
