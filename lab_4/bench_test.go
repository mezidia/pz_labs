package lab4

import (
	"fmt"
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/engine"
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/parser"
	"os"
	"testing"
)

var cmds []engine.Command

func BenchmarParser(b *testing.B) {
	l := 0
	for k := 0; k < 20; k++ {
		l += 2000
		f, _ := os.Create("test.txt")
		for j := 0; j < l; j++ {
			f.WriteString("print тутВсеОкВиводьМене\nreverse FaineMisto\n")
		}
		b.ResetTimer()
		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			cmds = parser.Parse("test.txt")
		})
	}
}
