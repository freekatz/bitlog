package common

// Pool a goroutine pool TODO
type Pool struct {
	poolSize int
}

type Runner func()

func (p *Pool) RunFunc(r func()) {
	p.Run(Runner(r))
}

func (p *Pool) Run(r Runner) {
	go r()
}
