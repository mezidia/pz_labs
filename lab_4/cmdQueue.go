package lab4

type cmdQueue struct {
	c []Command
}

func (q *cmdQueue) push(cmd Command) {
	q.c = append(q.c, cmd)
}

func (q *cmdQueue) pull() Command {
	res := q.c[0]
	q.c[0] = nil
	q.c = q.c[1:]
	return res
}
