package server

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/saicharith2012/goredis/internal/store"
)

type Server struct {
	port string
	store *store.SharedState
}

func New(port string, store *store.SharedState) *Server {
	return &Server{port: port, store: store}
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

	writer := bufio.NewWriter(conn)

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

		command := args[0]
		args = args[1:]
		fmt.Printf("parsed command and args from %v: %s, %#v\n", conn.RemoteAddr(), command, args)

		response := handleCommand(s , command, args)

		_, err = writer.WriteString(response)
		if err != nil {
			fmt.Printf("write error: %v\n", err)
			return
		}

		writer.Flush()

		fmt.Println("response sent to", conn.RemoteAddr())

	}
}
