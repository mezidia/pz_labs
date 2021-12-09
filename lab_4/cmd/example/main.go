package main

import (
	lab4 "github.com/mezidia/pz_labs/tree/lab4/lab_4"
)

func main() {
	l := new(lab4.Loop)
	l.Start()
	l.Post(lab4.PrintCommand{Arg: "asd"})

}
