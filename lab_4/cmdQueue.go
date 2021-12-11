package lab4

import (
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/engine"
)

type cmdQueue struct {
	c []engine.Command
}

func (q *cmdQueue) push(cmd engine.Command) {
	q.c = append(q.c, cmd)
}

func (q *cmdQueue) pull() engine.Command {
	res := q.c[0]
	q.c[0] = nil
	q.c = q.c[1:]
	return res
}

func (q *cmdQueue) size() int {
	return len(q.c)
}
