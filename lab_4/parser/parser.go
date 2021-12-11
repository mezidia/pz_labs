package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/engine"
)

func Parse(file string) ([]engine.Command, error) {
	input, err := os.Open(file)
	if err != nil {
		fmt.Print(err)
	}

	var cmds []engine.Command
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if len(strings.Trim(scanner.Text(), " ")) == 0 {
			continue
		}
		commands := strings.Split(scanner.Text(),"\n")
		for _, command := range commands{
			split := strings.Fields(command)
			if split[0] == "print" {
				if len(split) != 2 {
					err := fmt.Errorf("Too many arguments")
					return nil, err
				}
				cmd := &engine.PrintCommand{
					Arg: split[1],
				}
				cmds = append(cmds, cmd)
			} else if split[0] == "reverse" {
				if len(split) != 2 {
					err := fmt.Errorf("Too many arguments")
					return nil, err
				}
				cmd := &engine.ReverseCommand{
					Arg: split[1],
				}
				cmds = append(cmds, cmd)
			} else {
				err := fmt.Errorf("Command %s does not exist", split[0])
				return nil, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return nil, err
	}
	return cmds, nil
}
