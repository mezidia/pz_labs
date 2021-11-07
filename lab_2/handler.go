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

	text, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		text = strings.TrimRight(text, "\r\n")
	} else {
		text = strings.TrimRight(text, "\n")
	}
	result, err := PrefixToPostfix(text)
	if err != nil {
		return err
	} else {
		fmt.Fprintln(ch.Out, "postfix:", result)
	}
	return nil
}
