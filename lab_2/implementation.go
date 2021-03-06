package lab2

import (
	"fmt"
	"regexp"
	"strings"
)

func PrefixToPostfix(input string) (string, error) {
	//prep str
	// input = strings.ReplaceAll(input, " ", "")
	input = strings.Trim(input, " ")
	r := regexp.MustCompile("\\s+")
	input = r.ReplaceAllString(input, " ")
	arr := strings.Split(input, " ")
	fmt.Println(arr)
	l := len(arr)

	//prep stack, result var & regEx
	rExOperator := regexp.MustCompile(`(\+|-|\*|\/)`)
	rExOperand := regexp.MustCompile(`[0-9]`)
	st := Stack{}
	result := make([]string, 0)

	//main loop
	for i := l - 1; i >= 0; i-- {
		if rExOperand.MatchString(arr[i]) {
			//operand
			st.Push(arr[i])
		} else if rExOperator.MatchString(arr[i]) {
			//operator
			y, err := st.Pop()
			if err != nil {
				return "", fmt.Errorf("invalid expression")
			} else {
				result = append(result, y)
				if x, err := st.Pop(); err == nil {
					result = append(result, x)
				}
				result = append(result, arr[i])
			}
		} else {
			//bad char
			return "", fmt.Errorf("bad char")
		}
	}

	if len(result) <= 2 || st.Size() != 0 {
		return "", fmt.Errorf("invalid expression")
	}
	return strings.Join(result, " "), nil
}
