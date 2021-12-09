package parser

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func Parse(file string) {
	input, err := os.Open(file)
	if err != nil {
		fmt.Print(err)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		// cmd := parse(commandLine)
		
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// cmd := &printCommand{
	// 	arg: "Hello",
	// }


	// // Set the split function for the scanning operation.
	// scanner.Split(bufio.ScanWords)
	// // Count the words.
	// count := 0
	// for scanner.Scan() {
	// 	count++
	// }
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading input:", err)
	// }
	// fmt.Printf("%d\n", count)
}
