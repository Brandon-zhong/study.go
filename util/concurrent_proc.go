package util

import "sync"

type ConcurrentProc struct {
	size     int
	infoChan chan interface{}
	procFunc func(interface{})
	once     sync.Once
	wait sync.WaitGroup
}

func NewConcurrentProc(size int, procFunc func(interface{})) *ConcurrentProc {
	return &ConcurrentProc{
		size:     size,
		procFunc: procFunc,
		infoChan: make(chan interface{}, size),
	}
}

func (p *ConcurrentProc) Start() *ConcurrentProc {
	if p.size == 0 {
		panic("illegal size!")
	}
	p.once.Do(func() {
		p.infoChan = make(chan interface{}, p.size)
		for i := 0; i < p.size; i++ {
			go func() {
				p.wait.Add(1)
				defer p.wait.Done()
				for info := range p.infoChan {
					p.procFunc(info)
				}
			}()
		}
	})
	return p
}

func (p *ConcurrentProc) Wait() {
	p.wait.Wait()
}

func (p *ConcurrentProc) Done() {
	close(p.infoChan)
}

func (p *ConcurrentProc) AddInfo(info interface{}) {
	p.infoChan <- info
}
