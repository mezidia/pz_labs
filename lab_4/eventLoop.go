package lab4

type Command interface {
	Exec(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type Loop struct {
}

type cmdQueue struct {
}

func (q *cmdQueue) push(cmd Command) {

}

func (q *cmdQueue) pull(cmd Command) Command {
	return nil
}

func (l *Loop) Post(cmd Command) {
	cmd.Exec(l)
}

func (l *Loop) Start() {

}

func (l *Loop) AwaitStop() {

}

func main() {

}
