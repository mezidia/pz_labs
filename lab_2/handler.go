package lab2

import (
	"bufio"
	"fmt"
	"io"
	"runtime"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	// TODO: Add necessary fields.
	in  io.Reader
	out io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.in)
	fmt.Fprintln(ch.out, "Polish prefix notation to postfix")
	fmt.Fprintln(ch.out, "***************************************")

	for {
		fmt.Fprint(ch.out, "||# ")

		text, _ := reader.ReadString('\n')
		if runtime.GOOS == "windows" {
			text = strings.TrimRight(text, "\r\n")
		} else {
			text = strings.TrimRight(text, "\n")
		}
		result, err := PrefixToPostfix(text)
		if err != nil {
			fmt.Fprintln(ch.out, "Err converting:", err)
		} else {
			fmt.Fprintln(ch.out, "postfix:", result)
		}

	}
}
