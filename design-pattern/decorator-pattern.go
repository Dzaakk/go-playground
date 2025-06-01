package main

import (
	"log"
)

type Service interface {
	Execute() string
}

type BaseService struct{}

func (b *BaseService) Execute() string {
	return "Executing base service"
}

type LoggingDecorator struct {
	Wrapped Service
}

func (l *LoggingDecorator) Execute() string {
	log.Println("Before execution")
	res := l.Wrapped.Execute()
	log.Println("After execution")
	return res
}

// func main() {
// 	svc := &LoggingDecorator{Wrapped: &BaseService{}}
// 	fmt.Println(svc.Execute())
// }
