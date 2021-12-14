package engine

import (
  "fmt"
)

type PrintCommand struct { 
  Arg string 
} 

func (p *PrintCommand) Execute(loop Handler) { 
  fmt.Println(p.Arg) 
} 
