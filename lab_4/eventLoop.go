package lab4

import (
	"github.com/mezidia/pz_labs/tree/lab4/lab_4/engine"
)

type Loop struct {
	queue     *cmdQueue
	quit      chan bool
	isStopped bool
	isPaused  bool
}

func (l *Loop) Post(cmd engine.Command) {
	l.queue.push(cmd)
	if l.isPaused && !l.isStopped {
		l.startRoutine()
	}
}

func (l *Loop) startRoutine() {
	l.isPaused = false
	go func() {
		for {
			if l.queue.size() > 0 {
				cmd := l.queue.pull()
				cmd.Execute(l)
			} else if l.isStopped {
				l.quit <- true
				return
			} else {
				l.isPaused = true
				return
			}
		}
	}()
}

func (l *Loop) Start() {
	l.quit = make(chan bool, 1)
	l.queue = &cmdQueue{}
	l.startRoutine()
}

func (l *Loop) AwaitFinish() {
	l.isStopped = true
	<-l.quit
}
