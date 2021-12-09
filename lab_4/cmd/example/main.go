package main

import (
	lab4 "github.com/mezidia/pz_labs/tree/lab4/lab_4"
)

func main() {
	l := new(lab4.Loop)
	l.Start()
	l.Post(&lab4.PrintCommand{Arg: "asd"})
	l.Post(&lab4.PrintCommand{Arg: "asd1"})
	l.Post(&lab4.PrintCommand{Arg: "asd2"})
	l.Post(&lab4.PrintCommand{Arg: "asd3"})
	l.AwaitFinish()
	l.Post(&lab4.PrintCommand{Arg: "error"})
}
