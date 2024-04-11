package pool

import "os/exec"

type Command struct {
	ID   string
	name string
	arg  []string
}

type CommandResult struct {
	ID     string
	Result []byte
	err    error
}

type CommandExecPool struct {
	In   chan *Command
	Out  chan *CommandResult
	size int
}

func (p *CommandExecPool) Start() {
	for i := 0; i < p.size; i++ {
		go func() {
			for {
				select {
				case command := <-p.In:
					cmd := exec.Command(command.name, command.arg...)
					out, err := cmd.CombinedOutput()
					p.Out <- &CommandResult{
						ID:     command.ID,
						Result: out,
						err:    err,
					}
				}
			}
		}()
	}
}

func NewCommandExecPool(size int) *CommandExecPool {
	return &CommandExecPool{
		In:   make(chan *Command),
		Out:  make(chan *CommandResult),
		size: size,
	}
}
