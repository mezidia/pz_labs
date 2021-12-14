package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	lab2 "github.com/mezidia/pz_labs/tree/main/lab_2"
)

func main() {
	var eFlag = flag.String("e", "", "help message for flag e")
	var fFlag = flag.String("f", "", "help message for flag f")
	var oFlag = flag.String("o", "", "help message for flag o")
	var err error
	var handler lab2.ComputeHandler
	var isCorrectArgs bool
	isCorrectArgs = false
	flag.Parse()
	if *fFlag != "" && *eFlag == "" && *oFlag == "" {
		isCorrectArgs = true
		readfile, _ := os.Open(*fFlag)
		handler = lab2.ComputeHandler{
			In:  readfile,
			Out: os.Stdout,
		}
	}
	if *fFlag == "" && *eFlag != "" && *oFlag == "" {
		isCorrectArgs = true
		handler = lab2.ComputeHandler{
			In:  bytes.NewBufferString(*eFlag),
			Out: os.Stdout,
		}
	}
	if *fFlag != "" && *eFlag == "" && *oFlag != "" {
		isCorrectArgs = true
		readfile, _ := os.Open(*fFlag)
		writefile, _ := os.Open(*oFlag)
		handler = lab2.ComputeHandler{
			In:  readfile,
			Out: writefile,
		}
	}
	if *fFlag == "" && *eFlag != "" && *oFlag != "" {
		isCorrectArgs = true
		writefile, _ := os.Open(*oFlag)
		handler = lab2.ComputeHandler{
			In:  bytes.NewBufferString(*eFlag),
			Out: writefile,
		}
	}
	if isCorrectArgs == true {
		err = handler.Compute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Errors: %s", err)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Errors: wrong arguments")
	}
}
