package util

import (
	"context"
	"fmt"
)

type ProcNode struct {
	in    chan interface{}
	out   chan interface{}
	next  *ProcNode
	start bool
}

type Proc func(interface{}) interface{}

type SeqConcurrentProc struct {
	head     *ProcNode
	tail     *ProcNode
	fn       Proc
	ctx      context.Context
	CancelFn context.CancelFunc
}

func NewSeqConcurrentProc(size int, fn Proc) *SeqConcurrentProc {
	pool := new(SeqConcurrentProc)
	pool.fn = fn
	pool.ctx, pool.CancelFn = context.WithCancel(context.Background())
	for i := 0; i < size; i++ {
		//创建节点
		n := new(ProcNode)
		n.in = make(chan interface{})
		n.out = make(chan interface{})
		//起一个协程用来处理信息
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("something has error -->", err)
				}
				fmt.Println("goroutine exit!")
			}()
			for {
				select {
				case <-pool.ctx.Done():
					return
				case val := <-n.in:
					n.out <- pool.fn(val)
				}
			}
		}()
		if i == 0 {
			pool.head = n
			pool.tail = n
		}
		pool.tail.next = n
		pool.tail = n
	}
	return pool
}

func (p *SeqConcurrentProc) Process() (chan<- interface{}, <-chan interface{}) {
	in := make(chan interface{})
	out := make(chan interface{})
	go p.process(in, out)
	return in, out
}

func (p *SeqConcurrentProc) process(in <-chan interface{}, out chan<- interface{}) {
	for val := range in {
		//已经开始的节点出数据，没开始的节点开始处理并移动到队尾
		if p.head.start {
			out <- <-p.head.out
		} else {
			p.head.start = true
		}
		//移动节点，将头结点移动到尾节点并重新给任务信息
		cur := p.head
		p.head = p.head.next
		cur.next = nil
		p.tail.next = cur
		p.tail = cur
		p.tail.in <- val
	}
}
