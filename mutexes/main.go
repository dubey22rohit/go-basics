package main

import (
	"fmt"
	"sync"
)

type State struct {
	mu    sync.Mutex
	count int
}

// Maps are not concurrent safe, so we use mutexes
func (s *State) setState(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count = i
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
	wg.Wait() //* Blocks until wg are all done
	fmt.Printf("%+v", state)
}
