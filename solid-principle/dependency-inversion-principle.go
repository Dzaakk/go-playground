package main

import "fmt"

type Database interface {
	Save(data string)
}

type MySQL struct{}

func (m MySQL) Save(data string) {
	fmt.Println("SAve to MySQL")
}

type UserService struct {
	db Database
}

func (s UserService) Register(data string) {
	s.db.Save(data)
}

// ----------------------------
// wrong example
type PostgreSQL struct{}

func (p PostgreSQL) Save(data string) {
	fmt.Println("Save to MySQL")
}

type CustomerService struct {
	db PostgreSQL
}
