package lab4

import "fmt"

type Handler interface {
	Post(cmd Command)
}

type Command interface {
	Exec(handler Handler)
}

type PrintCommand struct {
	Arg string
}

func (p *PrintCommand) Exec(handler Handler) {
	fmt.Println(p.Arg)
}
