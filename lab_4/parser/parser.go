package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/engine"
)

func Parse(file string) []engine.Command {
	var cmds []engine.Command

	input, err := os.Open(file)
	if err != nil {
		err := fmt.Sprintf("INTERNAL ERROR: %s", err)
		cmd := &engine.PrintCommand{
			Arg: err,
		}
		cmds = append(cmds, cmd)
		return cmds
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if len(strings.Trim(scanner.Text(), " ")) == 0 {
			continue
		}
		commands := strings.Split(scanner.Text(),"\n")
		var cmd engine.Command
		for _, command := range commands{
			split := strings.Fields(command)
			if split[0] == "print" {
				if len(split) != 2 {
					err := fmt.Sprintf("SYNTAX ERROR on line *%s*", command)
					cmd = &engine.PrintCommand{
						Arg: err,
					}
				} else {
					cmd = &engine.PrintCommand{
						Arg: split[1],
					}
				}
			} else if split[0] == "reverse" {
				if len(split) != 2 {
					err := fmt.Sprintf("SYNTAX ERROR on line *%s*", command)
					cmd = &engine.PrintCommand{
						Arg: err,
					}
				} else {
					cmd = &engine.ReverseCommand{
						Arg: split[1],
					}
				}
			} else {
				err := fmt.Sprintf("Command *%s* does not exist", split[0])
				cmd = &engine.PrintCommand{
					Arg: err,
				}
			}
			cmds = append(cmds, cmd)
		}
	}

	if err := scanner.Err(); err != nil {
		err := fmt.Sprintf("INTERNAL ERROR: Reading standard input: %s", err)
		cmd := &engine.PrintCommand{
			Arg: err,
		}
		cmds = append(cmds, cmd)
	}
	return cmds
}
