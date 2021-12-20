package lab2

import (
	"fmt"
	"testing"
	"os"
	"strings"
)

var str string
var err error

func BenchmarkPrefixToPostfix(b *testing.B) {
	l := 0
	temp := os.Stdout
	for k := 0; k < 20; k++ {
		l += 100
		expr := strings.Repeat("+ 4 * - 4 2 3 ", l)
		b.ResetTimer()
		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			os.Stdout = nil
			str, err = PrefixToPostfix(expr)
			os.Stdout = temp
		})
	}
}