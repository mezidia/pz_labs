package parser

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	input, err := os.Open("./input")
	if err != nil {
		fmt.Print(err)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		cmd := parse(commandLine)
		// eventLoop.Post(cmd) 
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}


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
