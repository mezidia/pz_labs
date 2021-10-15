package lab2

import (
	"bufio"
	"fmt"
	"io"
	"runtime"
	"strings"
)

type ComputeHandler struct {
	In  io.Reader
	Out io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.In)
	fmt.Fprintln(ch.Out, "Polish prefix notation to postfix")
	fmt.Fprintln(ch.Out, "***************************************")

	for {
		fmt.Fprint(ch.Out, "||# ")

		text, _ := reader.ReadString('\n')
		if runtime.GOOS == "windows" {
			text = strings.TrimRight(text, "\r\n")
		} else {
			text = strings.TrimRight(text, "\n")
		}
		result, err := PrefixToPostfix(text)
		if err != nil {
			fmt.Fprintln(ch.Out, "Err converting:", err)
		} else {
			fmt.Fprintln(ch.Out, "postfix:", result)
		}

	}
}
