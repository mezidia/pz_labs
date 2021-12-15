package main

import (
	lab4 "github.com/mezidia/pz_labs/tree/lab4/lab_4"
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/engine"
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/parser"
)

func main() {
	cmds := parser.Parse("input")

	l := new(lab4.Loop)
	l.Start()
	for _, cmd := range cmds {
		l.Post(cmd)
	}
	l.AwaitFinish()
	l.Post(&engine.PrintCommand{Arg: "Should end!"})
}
