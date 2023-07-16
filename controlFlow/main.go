package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // *instead of a boolean channel, which takes 1-2 bytes
	// * we can make an empty struct which consumes 0 bytes
	msgch chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("server starting")
	s.loop() //block
}

func (s *Server) loop() {
	// * We can name the loops and now we can break outer for loop from inner select loop
mainloop:
	for {
		select {
		case ok := <-s.quitch:
			_ = ok
			fmt.Println("quitting server")
			break mainloop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
			//
		}
	}
	fmt.Println("server is shutting down gracefully")
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) handleMessage(msg string) {
	fmt.Print("we received a message: ", msg)
}

func (s *Server) quit() {
	// close(s.quitch) // *This will trigger the case in the select statement as well
	s.quitch <- struct{}{}
}

func main() {
	server := newServer()
	// server.start() //* This is blocking because of the loop,
	//* the program won't run below this line
	// go server.start() // * Now this is non-blocking, it is executed in its own thread

	// for i := 0; i < 10; i++ {
	// 	server.sendMessage(fmt.Sprintf("handle this number: %d\n", i))
	// }
	// time.Sleep(time.Second * 5)

	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	server.start()
}
