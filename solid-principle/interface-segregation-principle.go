package main

type Walker interface {
	Walk()
}

type Flyer interface {
	Fly()
}

type Dog struct{}

func (d Dog) Walk() {}

type Birds struct{}

func (b Birds) Walk() {}
func (b Birds) Fly()  {}

// ----------------------------
// wrong example
type Animal interface {
	Walk()
	Fly()
}

type Cat struct{}

func (c Cat) Walk() {}
func (c Cat) Fly()  {}
