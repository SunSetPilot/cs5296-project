package pool

import (
	"os/exec"
	"sync/atomic"
)

type Command struct {
	ID   string
	Name string
	Args []string
}

type CommandResult struct {
	ID     string
	Result []byte
	Err    error
}

type CommandExecPool struct {
	In      chan *Command
	Out     chan *CommandResult
	running int32
	size    int32
}

func (p *CommandExecPool) IsFull() bool {
	return atomic.LoadInt32(&p.running) == p.size
}

func (p *CommandExecPool) Start() {
	var i int32
	for i = 0; i < p.size; i++ {
		go func() {
			for {
				select {
				case command := <-p.In:
					atomic.AddInt32(&p.running, 1)
					cmd := exec.Command(command.Name, command.Args...)
					out, err := cmd.CombinedOutput()
					p.Out <- &CommandResult{
						ID:     command.ID,
						Result: out,
						Err:    err,
					}
					atomic.AddInt32(&p.running, -1)
				}
			}
		}()
	}
}

func NewCommandExecPool(size int32) *CommandExecPool {
	return &CommandExecPool{
		In:      make(chan *Command),
		Out:     make(chan *CommandResult),
		running: 0,
		size:    size,
	}
}
