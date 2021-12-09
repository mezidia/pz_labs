package lab4

type Loop struct {
	queue *cmdQueue
}

func (l *Loop) Post(cmd Command) {
	l.queue.push(cmd)
}

func (l *Loop) Start() {
	l.queue = &cmdQueue{}
	go func() {
		for {
			cmd := l.queue.pull()
			cmd.Exec(l)
		}
	}()
}

func (l *Loop) AwaitStop() {

}
