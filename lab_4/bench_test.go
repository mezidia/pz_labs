package lab4

import (
	"fmt"
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/engine"
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/parser"
	"math/rand"
	"os"
	"testing"
)

var temp = os.Stdout
var cmds []engine.Command

func BenchmarkEventLoop(b *testing.B) {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("error")
		return
	}
	b.SetParallelism(1)
	for k := 0; k < 10; k++ {
		l := 2*k + 2
		_, err := f.WriteString("print тутВсеОкВиводьМене\nreverse FaineMisto\n")
		if err != nil {
			fmt.Println("error")
			return
		}
		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			cmds = parser.Parse("test.txt")
		})
	}
}

func Count(s []int, a int) int {
	r := 0
	for _, e := range s {
		if a == e {
			r++
		}
	}
	return r
}

var cntRes int

func BenchmarkCount(b *testing.B) {
	const baseLen = 5
	for i := 0; i < 20; i++ {
		l := baseLen * 1 << i
		in := make([]int, l)
		for k := 0; k < l; k++ {
			in[k] = rand.Intn(10)
		}
		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			cntRes = Count(make([]int, l), 0)
		})
	}
}
