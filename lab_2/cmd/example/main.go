package main

import (
	"os"
	
	lab2 "github.com/mezidia/pz_labs/tree/main/lab_2"
)

func main() {

	handler := lab2.ComputeHandler{
		In:  os.Stdin,
		Out: os.Stdout,
	}
	handler.Compute()

}
