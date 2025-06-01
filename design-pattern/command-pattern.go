package main

import "fmt"

type Command interface {
	Execute()
}

type PrintCommand struct {
	Message string
}

func (p *PrintCommand) Execute() {
	fmt.Println(p.Message)
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) Run() {
	for _, cmd := range i.commands {
		cmd.Execute()
	}
}

// func main() {
// 	invoker := &Invoker{}

// 	cmd1 := &PrintCommand{Message: "Hello Dzak"}
// 	cmd2 := &PrintCommand{Message: "This is command pattern"}
// 	cmd3 := &PrintCommand{Message: "Bye"}

// 	invoker.AddCommand(cmd1)
// 	invoker.AddCommand(cmd2)
// 	invoker.AddCommand(cmd3)

// 	invoker.Run()
// }
