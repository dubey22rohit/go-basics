package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type State struct {
	count int32
}

func (s *State) setState(i int) {
	atomic.AddInt32(&s.count, int32(i))
}

func main() {
	state := &State{
		count: 0,
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			state.setState(i + 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v", state)
}
