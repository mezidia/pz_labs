package parser

import (
  "fmt"
  "github.com/mezidia/pz_labs/tree/lab4/lab_4/engine"
)

type printCommand struct { 
  arg string 
} 

func (p *printCommand) Execute(loop engine.Handler) { 
  fmt.Println(p.arg) 
} 
