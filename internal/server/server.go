package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type Server struct {
	port string
}

func New(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.port)

	if err != nil {
		return fmt.Errorf("listen failed: %s", err)
	}

	defer listener.Close()

	fmt.Println("goredis running on port", s.port)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("client connected", conn.RemoteAddr())

		go s.handleConnection(conn)

	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("client disconnected", conn.RemoteAddr())
	}()

	reader := bufio.NewReader(conn)

	for {

		args, err := respParser(reader)

		if err != nil {

			if err == io.EOF {
				fmt.Printf("client %s closing the connection...\n", conn.RemoteAddr())
			} else {
				fmt.Println("err:", err)
			}
			return
		}

		fmt.Printf("parsed command: %#v\n", args)
	}
}
