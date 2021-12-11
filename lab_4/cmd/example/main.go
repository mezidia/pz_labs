package main

import (
	"fmt"
	lab4 "github.com/mezidia/pz_labs/tree/lab4/lab_4"
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/parser"
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/engine"
)

func main() {
	cmds, err := parser.Parse("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	l := new(lab4.Loop)
	l.Start()
	for _, cmd := range(cmds) {
		l.Post(cmd)
	}
	l.AwaitFinish()	
	l.Post(&engine.PrintCommand{Arg: "Should end!"})
}
