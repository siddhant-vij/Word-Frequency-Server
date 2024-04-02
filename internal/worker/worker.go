package worker

import (
	"sync"

	"github.com/siddhant-vij/Word-Frequency-Server/pkg/common"
)

type Pool struct {
	Workers int
	Tasks   chan common.Task
	Results chan int
	Wg      sync.WaitGroup
}

func NewPool(workers int) *Pool {
	return &Pool{
		Workers: workers,
		Tasks:   make(chan common.Task, 100),
		Results: make(chan int, 100),
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.Workers; i++ {
		p.Wg.Add(1)
		go func(workerID int) {
			defer p.Wg.Done()
			for task := range p.Tasks {
				result := task.Execute()
				p.Results <- result
			}
		}(i)
	}

	go func() {
		p.Wg.Wait()
	}()
}
